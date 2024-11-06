package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"mail-service/app/config"
	"mail-service/utilities/model"
	"net/smtp"
)

func SetEmail(request []byte) error {
	var body bytes.Buffer
	var temp *template.Template
	var emailRequest model.EmailRequest
	conf, _ := config.Init()
	authorize := smtp.PlainAuth("", conf.Mail.Address, conf.Mail.Password, conf.Mail.Host)
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	_ = json.Unmarshal(request, &emailRequest)
	// Receiver
	var receiver = []string{
		emailRequest.Email,
	}

	// Set Email Body
	body.Write([]byte(fmt.Sprintf("Subject: Testing Send Email Kafka \n%s\n\n", mimeHeaders)))
	temp, _ = template.ParseFiles("template/template.html")

	temp.Execute(&body, "")
	err := smtp.SendMail(conf.Mail.Host+":"+conf.Mail.Port, authorize, conf.Mail.Address, receiver, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}
