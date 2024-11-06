package connections

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"

	"send-service/app/config"
	"send-service/pkg/broker"
)

var conf, _ = config.Init()

func GetKafkaConfig() *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0
	return kafkaConfig
}

func KafkaPublisherClient() (*broker.KafkaProducerAdapter, error) {
	kafkaConfig := GetKafkaConfig()
	producers, err := sarama.NewSyncProducer([]string{conf.Kafka.Endpoint}, kafkaConfig)
	if err != nil {
		return nil, err
	}
	logrus.Infof("Success create kafka sync-producer")

	kafka := &broker.KafkaProducerAdapter{
		Producer: producers,
	}
	return kafka, nil
}
