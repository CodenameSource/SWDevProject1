// main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	kafkamessenger "vot-hw1-checkprices/kafkaMessenger"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")

	kafkamessenger.SetupKafka(kafkaHost, kafkaPort)
	go kafkamessenger.StartRefreshEventConsumer()

	fmt.Println("Price checker service is running...")

	for {
		select {
		case <-signals:
			kafkamessenger.StopKafka()
			return
		}
	}
}
