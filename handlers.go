package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func fetchFromAPIFootball(w http.ResponseWriter, url string) {
	apiKey := os.Getenv("API_FOOTBALL_KEY")
	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Could not create request", http.StatusInternalServerError)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read API response", http.StatusInternalServerError)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, "Failed to parse API response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (cfg *apiConfig) getNextGame(w http.ResponseWriter, r *http.Request) {
	url := cfg.apiHost + "/fixtures?team=" + cfg.apiTeam + "&next=1"
	fetchFromAPIFootball(w, url)
}

func (cfg *apiConfig) getLastFiveGames(w http.ResponseWriter, r *http.Request) {
	url := cfg.apiHost + "/fixtures?team=" + cfg.apiTeam + "&last=5"
	fetchFromAPIFootball(w, url)
}
