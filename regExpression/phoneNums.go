package regExpression

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func ReadNumsFromFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Помилка при зчитуванні файлу:", err)
	}
	phoneNumbers := FindPhoneNumbers(string(data))
	for _, number := range phoneNumbers {
		fmt.Println("Номер телефону:", number)
	}
}

func FindPhoneNumbers(text string) []string {
	phonePattern := `(\d{3}-\d{3}-\d{4})|(\(\d{3}\)\s\d{3}-\d{4})|(\d{3}\s\d{3}\s\d{4})`
	reg := regexp.MustCompile(phonePattern)
	matches := reg.FindAllString(text, -1)
	return matches
}
