package main

import (
	"context"
	"encoding/json"
	"github.com/kaytu-io/kaytu-aws-describer/describer"
	config2 "github.com/kaytu-io/kaytu-util/pkg/config"
	"github.com/kaytu-io/kaytu-util/pkg/describe"
	"github.com/kaytu-io/kaytu-util/pkg/queue"
	"go.uber.org/zap"
)

func main() {
	j := DescriberJob{}
	panic(j.Run())
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
		var job describe.LambdaDescribeWorkerInput
		if err := json.Unmarshal(msg.Body, &job); err != nil {
			logger.Error("failed to consume message from LambdaDescribeWorkerInput", zap.Error(err))
			err = msg.Nack(false, false)
			if err != nil {
				logger.Error("failure while sending nack for message", zap.Error(err))
			}
			continue
		}

		err := describer.DescribeHandler(context.Background(), job)
		if err != nil {
			return err
		}
	}
	return nil
}
