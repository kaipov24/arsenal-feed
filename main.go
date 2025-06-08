package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	apiKey  string
	apiHost string
	apiTeam string
}

func main() {
	r := chi.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	apiCfg := apiConfig{
		apiKey:  os.Getenv("API_FOOTBALL_KEY"),
		apiHost: os.Getenv("API_FOOTBALL_API_HOST"),
		apiTeam: os.Getenv("API_FOOTBALL_API_TEAM"),
	}
	if apiCfg.apiKey == "" || apiCfg.apiHost == "" || apiCfg.apiTeam == "" {
		log.Fatal("API configuration is incomplete. Please set API_FOOTBALL_KEY, API_FOOTBALL_API_HOST, and API_FOOTBALL_API_TEAM in your .env file.")
	}
	r.Get("/next", apiCfg.getNextGame)
	r.Get("/last-five", apiCfg.getLastFiveGames)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
