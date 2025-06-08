package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
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
	url := cfg.apiHost + "/v3/fixtures?team=" + cfg.apiTeam + "&next=1"
	fetchFromAPIFootball(w, url)
}

func (cfg *apiConfig) getLastFiveGames(w http.ResponseWriter, r *http.Request) {
	url := cfg.apiHost + "/fixtures?team=" + cfg.apiTeam + "&last=5"
	fetchFromAPIFootball(w, url)
}

func (cfg *apiConfig) getCoach(w http.ResponseWriter, r *http.Request) {
	url := cfg.apiHost + "/coachs?team=" + cfg.apiTeam
	fetchFromAPIFootball(w, url)
}

func (cfg *apiConfig) getTeamInfo(w http.ResponseWriter, r *http.Request) {
	url := cfg.apiHost + "/teams?id=" + cfg.apiTeam
	fetchFromAPIFootball(w, url)
}

func (cfg *apiConfig) getLeagueStandings(w http.ResponseWriter, r *http.Request) {
	url := cfg.apiHost + "/standings?season=2024&team=" + cfg.apiTeam
	fetchFromAPIFootball(w, url)
}

func (cfg *apiConfig) getTransfers(w http.ResponseWriter, r *http.Request) {
	url := cfg.apiHost + "/transfers?team=" + cfg.apiTeam
	currentYear := time.Now().Year()

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

	filtered := []interface{}{}

	if response, ok := result["response"].([]interface{}); ok {
		for _, player := range response {
			playerMap, ok := player.(map[string]interface{})
			if !ok {
				continue
			}
			transfers, ok := playerMap["transfers"].([]interface{})
			if !ok {
				continue
			}
			newTransfers := []interface{}{}
			for _, t := range transfers {
				tMap, ok := t.(map[string]interface{})
				if !ok {
					continue
				}
				dateStr, ok := tMap["date"].(string)
				if !ok {
					continue
				}
				parsedDate, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					continue
				}
				year := parsedDate.Year()
				if year >= currentYear && year < currentYear+1 {
					newTransfers = append(newTransfers, tMap)
				}
			}
			if len(newTransfers) > 0 {
				playerMap["transfers"] = newTransfers
				filtered = append(filtered, playerMap)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"response": filtered,
	})

}
