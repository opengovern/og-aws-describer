package main

import (
	"context"
	"encoding/json"
	"github.com/opengovern/og-aws-describer/describer"
	config2 "github.com/opengovern/og-util/pkg/config"
	"github.com/opengovern/og-util/pkg/describe"
	"github.com/opengovern/og-util/pkg/queue"
	"go.uber.org/zap"
)

func main() {
	j := DescriberJob{}
	panic(j.Run())
}

type DescriberConfig struct {
	RabbitMQ  RabbitMQ
	QueueName string
}

type RabbitMQ struct {
	Service  string
	Username string
	Password string
}

type DescriberJob struct {
	config      DescriberConfig
	hopperQueue queue.Interface
}

func (h *DescriberJob) Run() error {
	config2.ReadFromEnv(&h.config, nil)

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	logger.Info("Initializing the scheduler")
	h.hopperQueue, err = initRabbitQueue(h.config.RabbitMQ, h.config.QueueName)
	if err != nil {
		return err
	}

	ch, err := h.hopperQueue.Consume()
	if err != nil {
		return err
	}

	for msg := range ch {
		var job describe.DescribeWorkerInput
		if err := json.Unmarshal(msg.Body, &job); err != nil {
			logger.Error("failed to consume message from DescribeWorkerInput", zap.Error(err))
			err = msg.Nack(false, false)
			if err != nil {
				logger.Error("failure while sending nack for message", zap.Error(err))
			}
			continue
		}

		err := describer.DescribeHandler(context.Background(), describer.TriggeredByAWSLambda, job)
		if err != nil {
			logger.Error("failure while handling job", zap.Uint("jobID", job.DescribeJob.JobID))
			continue
		}
	}
	return nil
}

func initRabbitQueue(cnf RabbitMQ, queueName string) (queue.Interface, error) {
	qCfg := queue.Config{}
	qCfg.Server.Username = cnf.Username
	qCfg.Server.Password = cnf.Password
	qCfg.Server.Host = cnf.Service
	qCfg.Server.Port = 5672
	qCfg.Queue.Name = queueName
	qCfg.Queue.Durable = true
	qCfg.Producer.ID = "describe-scheduler"
	insightQueue, err := queue.New(qCfg)
	if err != nil {
		return nil, err
	}

	return insightQueue, nil
}
