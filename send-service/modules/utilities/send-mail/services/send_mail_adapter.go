package services

import (
	"send-service/pkg/broker"
)

type Service interface {
	SendMail(mail string) error
}

type service struct {
	Kafka *broker.KafkaProducerAdapter
}

func NewService(kafka *broker.KafkaProducerAdapter) *service {
	return &service{kafka}
}
