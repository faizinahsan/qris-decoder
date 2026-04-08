package main

import (
	"log"

	httpserver "faizinahsan/qris-decoder/interfaces/http"
)

func main() {
	r := httpserver.NewRouter()
	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
