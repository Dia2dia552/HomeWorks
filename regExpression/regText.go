package regExpression

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func ReadTextFromFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Помилка при зчитуванні файлу:", err)
	}
	text := string(data)
	matchingWords := FindWordsMatchingPattern(text)
	for _, word := range matchingWords {
		fmt.Println("Слово:", word)
	}
}
func FindWordsMatchingPattern(text string) []string {
	pattern := `(?i)\b[аеіоу][а-я]*[^аеіоу]\b`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)
	return matches
}
