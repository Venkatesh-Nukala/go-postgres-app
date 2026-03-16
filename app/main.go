package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {

	connStr := "host=db port=5432 user=postgres password=postgres dbname=test sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(w, "Database connection error")
		return
	}

	defer db.Close()

	fmt.Fprintf(w, "Hello from Go App connected to PostgreSQL!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
