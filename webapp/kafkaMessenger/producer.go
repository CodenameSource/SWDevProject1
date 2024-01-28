package kafkamessenger

import (
	"fmt"

	"github.com/IBM/sarama"
)

var Producer sarama.SyncProducer

const refreshTopic = "refresh-price"

func SendRefreshEvent(url string) error {
	fmt.Println("Sending refresh event with url: " + url)

	msg := sarama.ProducerMessage{
		Topic: refreshTopic,
		Key:   sarama.StringEncoder(url),
		Value: sarama.StringEncoder("True"),
	}

	_, _, err := Producer.SendMessage(&msg)
	return err
}
