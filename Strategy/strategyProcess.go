package Strategy

import (
	"fmt"
	"regexp"
	"strings"
)

type TextProcessingStrategy interface {
	Process(text string) string
}
type WordCountStrategy struct{}

func (wcs WordCountStrategy) Process(text string) string {
	words := strings.Fields(text)
	count := len(words)
	return fmt.Sprintf("Кількість слів: %d", count)
}

type ReplaceSpacesDecorator struct{}

func (rsd ReplaceSpacesDecorator) Process(text string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(text, " ")
}

type RemoveHTMLTagsDecorator struct{}

func (rtd RemoveHTMLTagsDecorator) Process(text string) string {
	re := regexp.MustCompile(`<[^>]+>`)
	return re.ReplaceAllString(text, "")
}

type TextProcessor struct {
	strategies []TextProcessingStrategy
}

func NewTextProcessor(strategies ...TextProcessingStrategy) TextProcessor {
	return TextProcessor{
		strategies: strategies,
	}
}
func (tp TextProcessor) Process(text string) string {
	result := text
	for _, strategy := range tp.strategies {
		result = strategy.Process(result)
	}
	return result
}
