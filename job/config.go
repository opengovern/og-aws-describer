package main

type DescriberConfig struct {
	RabbitMQ  RabbitMQ
	QueueName string
}

type RabbitMQ struct {
	Service  string
	Username string
	Password string
}
