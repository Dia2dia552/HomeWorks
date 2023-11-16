package main

import (
	"NewProject/routes"
	"net/http"
)

func main() {
	routes.SetupRoutes()

	port := ":8080"
	http.ListenAndServe(port, nil)
}
