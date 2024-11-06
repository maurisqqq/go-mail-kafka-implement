package connections

import (
	"os"
	"time"

	"mail-service/app/config"
	"mail-service/pkg/broker"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func GetKafkaConfig() *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	return kafkaConfig
}

var conf, _ = config.Init()

func KafkaConsumerClient() {
	kafkaConfig := GetKafkaConfig()
	consumers, err := sarama.NewConsumer([]string{conf.Kafka.Endpoint}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Error create kakfa consumer got error %v", err)
	}
	defer func() {
		if err := consumers.Close(); err != nil {
			panic(err)
		}
	}()
	//Kafka consumer
	kafkaConsumer := &broker.KafkaConsumer{
		Consumer: consumers,
	}
	signals := make(chan os.Signal, 1)

	listTopics := []string{"send_email"}
	kafkaConsumer.Consume(listTopics, signals)
}
