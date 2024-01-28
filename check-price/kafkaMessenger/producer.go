package kafkamessenger

import (
	"strconv"

	"github.com/IBM/sarama"
)

var Producer sarama.SyncProducer

const priceTopic = "refreshed-price"

func SendRefreshedPrice(url string, price float64) error {
	priceStr := strconv.FormatFloat(price, 'f', -1, 64)

	message := &sarama.ProducerMessage{
		Topic: priceTopic,
		Key:   sarama.StringEncoder(url),
		Value: sarama.StringEncoder(priceStr),
	}

	_, _, err = Producer.SendMessage(message)
	return err
}
