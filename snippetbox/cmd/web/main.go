package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "root:vivas@123 sslmode=disable", "Postgres Datasource Name")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func openDB(dsn string) *sql.DB {

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db

}
