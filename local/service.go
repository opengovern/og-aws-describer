package local

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/kaytu-io/kaytu-aws-describer/describer"
	"github.com/kaytu-io/kaytu-util/pkg/config"
	"github.com/kaytu-io/kaytu-util/pkg/describe"
	esSinkClient "github.com/kaytu-io/kaytu-util/pkg/es/ingest/client"
	"github.com/kaytu-io/kaytu-util/pkg/jq"
	"github.com/kaytu-io/kaytu-util/pkg/kaytu-es-sdk"
	"github.com/kaytu-io/kaytu-util/pkg/koanf"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"runtime"
	"time"
)

const (
	StreamName    = "kaytu_aws_describer"
	JobQueueTopic = "kaytu_aws_describer_job_queue"
	ConsumerGroup = "aws-describer"
)

type Config struct {
	NATS config.NATS `koanf:"nats"`
}

func WorkerCommand() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cmd.SilenceUsage = true
			cnf := koanf.Provide("aws_describer", Config{})
			logger, err := zap.NewProduction()
			if err != nil {
				return err
			}

			w, err := NewWorker(
				cnf,
				logger,
				cmd.Context(),
			)
			if err != nil {
				return err
			}

			return w.Run(ctx)
		},
	}

	return cmd
}

type Worker struct {
	config   Config
	logger   *zap.Logger
	esClient kaytu.Client
	jq       *jq.JobQueue

	esSinkClient esSinkClient.EsSinkServiceClient
}

func NewWorker(
	config Config,
	logger *zap.Logger,
	ctx context.Context,
) (*Worker, error) {
	jq, err := jq.New(config.NATS.URL, logger)
	if err != nil {
		logger.Error("failed to create job queue", zap.Error(err), zap.String("url", config.NATS.URL))
		return nil, err
	}

	if err := jq.Stream(ctx, StreamName, "aws describe job runner queue", []string{JobQueueTopic}, 200000); err != nil {
		logger.Error("failed to create stream", zap.Error(err))
		return nil, err
	}

	w := &Worker{
		config: config,
		logger: logger,
		jq:     jq,
	}

	return w, nil
}

func (w *Worker) Run(ctx context.Context) error {
	w.logger.Info("starting to consume")

	consumeCtx, err := w.jq.ConsumeWithConfig(ctx, ConsumerGroup, StreamName, []string{JobQueueTopic}, jetstream.ConsumerConfig{
		Replicas:          1,
		AckPolicy:         jetstream.AckExplicitPolicy,
		DeliverPolicy:     jetstream.DeliverAllPolicy,
		MaxAckPending:     -1,
		AckWait:           time.Minute * 30,
		InactiveThreshold: time.Hour,
	}, []jetstream.PullConsumeOpt{
		jetstream.PullMaxMessages(1),
	}, func(msg jetstream.Msg) {
		w.logger.Info("received a new job")

		defer msg.Ack()

		ctx, cancel := context.WithTimeoutCause(ctx, time.Minute*25, errors.New("describe worker timed out"))
		defer cancel()

		if err := w.ProcessMessage(ctx, msg); err != nil {
			w.logger.Error("failed to process message", zap.Error(err))
		}
		err := msg.Ack()
		if err != nil {
			w.logger.Error("failed to ack message", zap.Error(err))
		}

		w.logger.Info("processing a job completed")
	})
	if err != nil {
		return err
	}

	w.logger.Info("consuming")

	<-ctx.Done()
	consumeCtx.Drain()
	consumeCtx.Stop()

	return nil
}

func (w *Worker) ProcessMessage(ctx context.Context, msg jetstream.Msg) error {
	startTime := time.Now()
	var input describe.DescribeWorkerInput
	err := json.Unmarshal(msg.Data(), &input)
	if err != nil {
		return err
	}
	runtime.GC()

	w.logger.Info("running job", zap.Uint("id", input.DescribeJob.JobID), zap.String("type", input.DescribeJob.ResourceType), zap.String("account", input.DescribeJob.AccountID))

	err = describer.DescribeHandler(ctx, w.logger, describer.TriggeredByLocal, input)
	endTime := time.Now()

	w.logger.Info("job completed", zap.Uint("id", input.DescribeJob.JobID), zap.String("type", input.DescribeJob.ResourceType), zap.String("account", input.DescribeJob.AccountID), zap.Duration("duration", endTime.Sub(startTime)))
	if err != nil {
		w.logger.Error("failure while running job", zap.Error(err))
		return err
	}

	return nil
}
