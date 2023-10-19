package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TranslationRequest struct {
	Text       string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}

type TranslationResponse struct {
	TranslatedText string `json:"translated_text"`
}

func main() {
	http.HandleFunc("/translate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
			return
		}
		var request TranslationRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		translatedText, err := translateText(request.Text, request.SourceLang, request.TargetLang)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := TranslationResponse{TranslatedText: translatedText}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func translateText(text, sourceLang, targetLang string) (string, error) {
	apiKey := "API_KEY"
	url := fmt.Sprintf("https://translation.googleapis.com/language/translate/v2?key=%s", apiKey)
	data := fmt.Sprintf(`{"q":"%s","source":"%s","target":"%s","format":"text"}`, text, sourceLang, targetLang)
	resp, err := http.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var translationResponse map[string]interface{}
	err = json.Unmarshal(body, &translationResponse)
	if err != nil {
		return "", err
	}
	translatedText := translationResponse["data"].(map[string]interface{})["translations"].([]interface{})[0].(map[string]interface{})["translatedText"].(string)
	return translatedText, nil
}
