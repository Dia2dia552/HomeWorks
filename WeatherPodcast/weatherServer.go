package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func GetWeatherData(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "Параметр 'city' не вказаний", http.StatusBadRequest)
		return
	}

	apiKey := "e5408c8173744f9c99694825231910"
	apiUrl := "https://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + city

	response, err := http.Get(apiUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		http.Error(w, "Не вдалось отримати дані з API", http.StatusBadGateway)
		return
	}

	var weatherData map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(weatherData)
	if err != nil {
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather", GetWeatherData).Methods("GET")
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
