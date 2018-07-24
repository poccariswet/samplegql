package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func NewConfig() error {
	db, err := sqlx.Connect("postgres", "user=ql_user password=graphql dbname=sampledb sslmode=disable")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func main() {
	if err := NewConfig(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if err := DB.Ping(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("connect!")
	}

	http.HandleFunc("/", QLhandler)
	http.ListenAndServe(":8080", nil)
}
