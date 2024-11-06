package main

import (
	"mail-service/app/config"
	"mail-service/app/connections"
	"mail-service/pkg/logs"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	var conf, _ = config.Init()

	// Listen to kafka
	go connections.KafkaConsumerClient()

	// Format loging
	logs.SetFormatter()

	logrus.Infof("🚀 Succeed Running On http://localhost%v", conf.App.Port)

	http.ListenAndServe(conf.App.Port, nil)
}
