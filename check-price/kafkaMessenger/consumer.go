package kafkamessenger

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

const refreshTopic = "refresh-price"

var (
	Consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
)

func StartRefreshEventConsumer() {
	partitionConsumer, err = Consumer.ConsumePartition(refreshTopic, 0, sarama.OffsetNewest) // Replace with your Kafka topic
	if err != nil {
		log.Fatal("Error creating Kafka partition consumer:", err)
	}

	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Println("Received event from Kafka: ", string(msg.Key), ":", string(msg.Value))

			if string(msg.Value) != "True" {
				continue
			}

			url := string(msg.Key)

			ExtractPrice(url)
			continue
		case <-StopConsumer:
			return
		}
	}
}
