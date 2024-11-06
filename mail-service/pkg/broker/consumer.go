package broker

import (
	"mail-service/utilities/service"
	"os"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	Consumer sarama.Consumer
}

// Consume function to consume message from apache kafka
func (c *KafkaConsumer) Consume(topics []string, signals chan os.Signal) {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	for _, topic := range topics {
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("Unable to get partition got error %v", err)
		}
		for _, partition := range partitionList {
			go consumeMessage(c.Consumer, topic, partition, chanMessage)
		}
	}

	logrus.Infof("Kafka is consuming....")

	for {
		select {
		case msg := <-chanMessage:
			if msg.Topic == "send_email" {
				// hit input_harga_storage
				logrus.Infof("Consume and redirect to send email")
				err := service.SetEmail([]byte(msg.Value))
				if err != nil {
					logrus.Errorf("error %v", err)
				}
			}
		case sig := <-signals:
			if sig == os.Interrupt {
				logrus.Errorf("Signal is interupted")
			}
		}
	}
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		logrus.Errorf("Unable to consume partition got error %v", err)
	}

	defer func() {
		if err := msg.Close(); err != nil {
			logrus.Errorf("Unable to close partition : %v", err)
		}
	}()

	for {
		msg := <-msg.Messages()
		c <- msg
	}
}
