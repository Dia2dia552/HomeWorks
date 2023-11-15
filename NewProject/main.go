package main

import (
	"NewProject/routes"
	"net/http"
)

func main() {
	routes.SetupRoutes()

	port := ":8080"

	// Запуск сервера
	http.ListenAndServe(port, nil)
}
