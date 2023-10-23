package main

import (
	"Strategy/Strategy"
	"fmt"
	"os"
)

func main() {
	//regExpression.ReadNumsFromFile("contacts.txt")
	//regExpression.ReadTextFromFile("text.txt")

	fileContents, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Помилка зчитування файлу:", err)
		return
	}
	text := string(fileContents)

	fmt.Println("Виберіть алгоритм обробки:")
	fmt.Println("1. Підрахунок слів")
	fmt.Println("2. Заміна подвійних пробілів та видалення HTML тегів")

	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil {
		return
	}

	var processor Strategy.TextProcessor

	switch choice {
	case 1:
		processor = Strategy.NewTextProcessor(Strategy.WordCountStrategy{})
	case 2:
		processor = Strategy.NewTextProcessor(
			Strategy.ReplaceSpacesDecorator{},
			Strategy.RemoveHTMLTagsDecorator{},
		)
	default:
		fmt.Println("Неправильний вибір алгоритму")
		return
	}

	result := processor.Process(text)
	fmt.Println("Результат обробки:")
	fmt.Println(result)
	err = os.WriteFile("output.txt", []byte(result), 0644)
	if err != nil {
		fmt.Println("Помилка збереження результату:", err)
	}
}
