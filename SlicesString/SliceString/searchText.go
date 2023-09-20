package SliceString

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchEditorText(editorText []string, searchText string) []string {
	var results []string
	for _, line := range editorText {
		if strings.Contains(line, searchText) {
			results = append(results, line)
		}
	}
	return results
}
func StartSearch() {
	editorText := []string{
		"перший рядок",
		"другий рядок",
		"третій рядок",
		"четвертий рядок",
	}
	fmt.Println("Введіть текст для пошуку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchText := scanner.Text()
	results := (searchEditorText(editorText, searchText))

	fmt.Printf("Результати пошуку: '%s': /n", searchText)
	for _, result := range results {
		fmt.Println(result)
	}
}
