package producer

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

type Orange struct {
	Size int `json:"size"`
}

func generateRandomSize() int {
	return rand.Intn(20) + 5
}

func sendOrangesToQueue(ch *amqp.Channel, queueName string) {
	_, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	for i := 0; i < 10; i++ {
		orange := Orange{Size: generateRandomSize()}
		body, err := json.Marshal(orange)
		if err != nil {
			log.Printf("Failed to marshal JSON: %s", err)
			continue
		}

		err = ch.Publish(
			"",
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			log.Printf("Failed to publish a message: %s", err)
			continue
		}

		fmt.Printf("Sent an orange with size: %d cm\n", orange.Size)
		time.Sleep(1 * time.Second)
	}
}

func StartSending() {
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
	go sendOrangesToQueue(ch, queueName)
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
