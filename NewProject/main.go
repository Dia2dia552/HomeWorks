package main

import (
	"NewProject/routes"
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "user=username password=password dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	routes.SetupRoutes()

	port := ":8080"
	http.ListenAndServe(port, nil)
}
