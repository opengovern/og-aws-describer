package main

import "github.com/kaytu-io/kaytu-util/pkg/queue"

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
