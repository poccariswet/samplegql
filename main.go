package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DB *sqlx.DB
}

func NewConfig() (*Config, error) {
	db, err := sqlx.Connect("postgres", "user=ql_user password=graphql dbname=sampledb sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: db,
	}, nil
}

func main() {
	c, err := NewConfig()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if err := c.DB.Ping(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("connect!")
	}

	http.HandleFunc("/", QLhandler)
	http.ListenAndServe(":8080", nil)
}
