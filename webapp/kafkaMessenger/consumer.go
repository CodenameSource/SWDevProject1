package kafkamessenger

import (
	"fmt"
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

const priceTopic = "refreshed-price"

var (
	Consumer          sarama.Consumer
	PartitionConsumer sarama.PartitionConsumer
)

func StartKafkaConsumerPrice() {

	PartitionConsumer, err = Consumer.ConsumePartition(priceTopic, 0, sarama.OffsetNewest) // Replace with your Kafka topic
	if err != nil {
		log.Fatal("Error creating Kafka partition consumer:", err)
	}
}

func ListenForRefreshPrice() (string, float64, error) {
	for {
		select {
		case msg := <-PartitionConsumer.Messages():
			fmt.Printf("Received active message: %s\n", string(msg.Value))
			if msg.Key == nil || msg.Value == nil {
				continue
			}
			price, err := strconv.ParseFloat(string(msg.Value), 64)

			fmt.Println("Returning from ListenForRefreshPrice")
			return string(msg.Key), price, err
		case <-StopConsumer:
			fmt.Println("Stopping consumer")
			return "", 0, nil
		}
	}
}
