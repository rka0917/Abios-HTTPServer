package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rka0917/Abios-HTTPServer/models"
)

// Handler used for handling the /players endpoint
// Fetches all live series, checks the player roster and then fetches player information
// for players in the active roster.
func LivePlayersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get all live series
	liveSeries, err := models.GetLiveSeries()

	if err != nil {
		http.Error(w, "Error occurred while retrieving data.", http.StatusInternalServerError)
		return
	}

	liveSeriesIds := make([]int, len(liveSeries))

	for i, series := range liveSeries {
		liveSeriesIds[i] = int(series["id"].(float64))
	}

	// Get all the players for each series
	seriesRosters, err := models.GetSeriesRosters(liveSeriesIds)

	if err != nil {
		http.Error(w, "Error occurred while retrieving data.", http.StatusInternalServerError)
		return
	}

	playerIds := extractPlayerIdsFromRoster(seriesRosters)

	players, err := models.GetPlayersFromIds(playerIds)

	if err != nil {
		http.Error(w, "Error occurred while retrieving data.", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(players, "", "  ")

	if err != nil {
		http.Error(w, "Error occurred while retrieving data.", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonData))
}

func extractPlayerIdsFromRoster(seriesRosters []map[string]interface{}) []int {
	playerIds := []int{}
	seenPlayers := make(map[int]bool)

	for _, roster := range seriesRosters {

		lineUp, ok := roster["line_up"].(map[string]interface{})
		if ok {
			players, ok := lineUp["players"].([]interface{})
			if ok {
				for _, player := range players {
					playerMap, ok := player.(map[string]interface{})
					if ok {
						playerId := int(playerMap["id"].(float64))
						_, playerAdded := seenPlayers[playerId]
						if !playerAdded {
							seenPlayers[playerId] = true
							playerIds = append(playerIds, playerId)
						}
					}
				}
			}
		}
	}
	return playerIds
}
