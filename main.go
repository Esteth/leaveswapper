package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esteth/leaveswapper/app"
)

func main() {
	port := os.Getenv("PORT")
	dburl := os.Getenv("DATABASE_URL")

	log.Println(port, dburl)

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if dburl == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	err := app.ListenAndServe(":"+port, dburl)

	if err != nil {
		fmt.Printf(err.Error())
	}
}
