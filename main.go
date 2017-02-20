package main

import "github.com/esteth/leaveswapper/app"
import "os"
import "log"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app.ListenAndServe(":" + port)
}
