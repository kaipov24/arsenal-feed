package main

import (
	"encoding/json"
	"io"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	http.HandleFunc("/last", getLastGame)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getLastGame(w http.ResponseWriter, r *http.Request) {
	url := "https://api-football-v1.p.rapidapi.com/v3/fixtures?team=42&next=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Could not create request", http.StatusInternalServerError)
		return
	}

	apiKey := os.Getenv("API_FOOTBALL_KEY")

	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		return
	}
	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "API request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, "Failed to parse API response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
