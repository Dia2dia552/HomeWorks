package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Orange struct {
	Size int `json:"size"`
}

func processOranges(ch *amqp.Channel, queueName string) {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	sizes := make(map[string]int)

	for msg := range msgs {
		var orange Orange
		err := json.Unmarshal(msg.Body, &orange)
		if err != nil {
			log.Printf("Failed to unmarshal JSON: %s", err)
			continue
		}

		sizeCategory := getSizeCategory(orange.Size)
		sizes[sizeCategory]++
		fmt.Printf("Received an orange with size: %d cm\n", orange.Size)
	}

	go func() {
		for {
			printSizeCounts(sizes)
			time.Sleep(60 * time.Second)
		}
	}()
}

func getSizeCategory(size int) string {
	if size <= 10 {
		return "Small"
	} else if size <= 20 {
		return "Medium"
	}
	return "Large"
}

func printSizeCounts(sizes map[string]int) {
	for category, count := range sizes {
		fmt.Printf("%s: %d\n", category, count)
	}
	fmt.Println("------------------")
}

func StartProcessing() {
	conn, ch := establishConnection()
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	queueName := "orange_queue"
	processOranges(ch, queueName)
}

func establishConnection() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	return conn, ch
}
