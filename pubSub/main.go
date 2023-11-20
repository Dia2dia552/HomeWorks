package main

import (
	"fmt"
	"pubSub/pubSub"
)

func main() {
	dirPath := "/HomeWorks"

	publisher := pubSub.NewPublisher()

	subscriber := make(chan string)
	publisher.Subscribe(subscriber)

	go func() {
		for {
			select {
			case event := <-subscriber:
				fmt.Println("Received notification:", event)
			}
		}
	}()

	pubSub.StartMonitoring(dirPath, publisher)
	select {}
}
