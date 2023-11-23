package main

import (
	"github.com/Dia2dia552/fruit-sorter/consumer"
	"github.com/Dia2dia552/fruit-sorter/producer"
)

func main() {
	producer.StartSending()
	consumer.StartProcessing()
	select {}
}
