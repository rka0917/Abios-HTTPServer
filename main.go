package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rka0917/Abios-HTTPServer/handlers"
)

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/series/live", handlers.LiveSeriesHandler)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
