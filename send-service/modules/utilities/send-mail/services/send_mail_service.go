package services

import (
	"encoding/json"
	sendMailModels "send-service/modules/utilities/send-mail/models"
)

func (s *service) SendMail(mail string) error {
	// Publish to kafka
	data, err := json.Marshal(sendMailModels.EmailRequest{
		Email: mail,
	})

	if err != nil {
		return err
	}

	err = s.Kafka.SendData("send_email", data)
	if err != nil {
		return err
	}

	// Success
	return nil
}
