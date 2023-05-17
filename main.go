package main

import (
	"log"
	"net/http"

	"github.com/didip/tollbooth/v7"
	"github.com/joho/godotenv"
	"github.com/rka0917/Abios-HTTPServer/handlers"
)

func main() {
	godotenv.Load(".env")

	// For now, we will go with fixed request limiting of 5 requests/s.
	lmt := tollbooth.NewLimiter(5, nil)
	lmt.SetMessage("You have reached maximum request limit")
	http.Handle("/series/live", tollbooth.LimitFuncHandler(lmt, handlers.LiveSeriesHandler))
	http.Handle("/players/live", tollbooth.LimitFuncHandler(lmt, handlers.LivePlayersHandler))
	http.Handle("/teams/live", tollbooth.LimitFuncHandler(lmt, handlers.LiveTeamsHandler))

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
