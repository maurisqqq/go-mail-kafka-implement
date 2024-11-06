package broker

import "github.com/IBM/sarama"

type KafkaProducerAdapter struct {
	Producer sarama.SyncProducer
}
