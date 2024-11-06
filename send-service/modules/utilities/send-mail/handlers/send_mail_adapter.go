package handlers

import (
	sendMailService "send-service/modules/utilities/send-mail/services"
	"send-service/pkg/broker"
)

type SendMailHandler struct {
	SendMailService sendMailService.Service
}

func NewSendMailHandler(kafka *broker.KafkaProducerAdapter) *SendMailHandler {
	sendMailService := sendMailService.NewService(kafka)

	return &SendMailHandler{sendMailService}
}
