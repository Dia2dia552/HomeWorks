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

	users, err := publisher.GetDataFromDB()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
	pubSub.StartMonitoring(dirPath, publisher)
	select {}
}
