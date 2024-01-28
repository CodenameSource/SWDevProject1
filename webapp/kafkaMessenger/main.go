package kafkamessenger

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

var (
	StopConsumer chan bool
	err          error
)

func SetupKafka(kafkaHost string, kafkaPort string) {
	brokers := []string{kafkaHost + ":" + kafkaPort} // Replace with your Kafka broker address

	StopConsumer = make(chan bool)

	// Setup Kafka producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Retry.Backoff = 1000 * time.Millisecond
	config.Producer.Return.Successes = true

	Producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal("Error creating Kafka producer:", err)
	}

	// Setup Kafka consumer
	config.Consumer.Return.Errors = true
	Consumer, err = sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Fatal("Error creating Kafka consumer:", err)
	}
}

func StopKafka() {
	StopConsumer <- true

	PartitionConsumer.Close()

	err := Consumer.Close()

	if err != nil {
		log.Fatal("Error closing Kafka consumer:", err)
	}

	Producer.Close()

	if err != nil {
		log.Fatal("Error closing Kafka producer:", err)
	}
}
