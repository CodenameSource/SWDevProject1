// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"webapp/backend"
	kafkamessenger "webapp/kafkaMessenger"

	"github.com/gorilla/handlers"
)


func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")

	kafkamessenger.SetupKafka(kafkaHost, kafkaPort)

	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDB := os.Getenv("MYSQL_DB")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDB + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:password@tcp(localhost:" + dbPort + ")/vothw1?charset=utf8mb4&parseTime=True&loc=Local"
	router, err := backend.Init(dsn)

	// Enable CORS for all routes
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Use handlers.CORS to enable CORS with the allowed headers, methods, and origins
	http.Handle("/", handlers.CORS(headers, methods, origins)(router))

	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	go kafkamessenger.StartKafkaConsumerPrice()

	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	for {
		select {
		case <-signals:
			kafkamessenger.StopKafka()
			return
		}
	}
}
