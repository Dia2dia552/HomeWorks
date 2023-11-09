package main

import (
	"Calculator/calculator"
	"fmt"
	"log"
)

func main() {
	expression := "3 + 5 * 2 - 8 / 4"
	result, err := calculator.Calculate(expression)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Результат: %f\n", result)
}
