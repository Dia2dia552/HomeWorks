package main

import (
	"sorter/consumer"
	producer "sorter/produser"
)

func main() {
	producer.StartSending()
	consumer.StartProcessing()
	select {}
}
