package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rka0917/Abios-HTTPServer/models"
)

// Handler used to handle /teams endpoint.
// Fetches all live series, checks the participating teams and then fetches detailed team information.
func LiveTeamsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	liveSeries, err := models.GetLiveSeries()

	if err != nil {
		http.Error(w, "Error occurred while retrieving team data.", http.StatusInternalServerError)
		return
	}

	liveSeriesIds := make([]int, len(liveSeries))

	for i, series := range liveSeries {
		liveSeriesIds[i] = int(series["id"].(float64))
	}

	// Get all the players for each series
	seriesRosters, err := models.GetSeriesRosters(liveSeriesIds)

	if err != nil {
		http.Error(w, "Error occurred while retrieving team data.", http.StatusInternalServerError)
		return
	}

	teamIds := extractTeamIdsFromRoster(seriesRosters)

	players, err := models.GetTeamsFromIds(teamIds)

	if err != nil {
		http.Error(w, "Error occurred while retrieving team data.", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(players, "", "  ")

	if err != nil {
		http.Error(w, "Error occurred while retrieving team data.", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonData))
}

func extractTeamIdsFromRoster(seriesRosters []map[string]interface{}) []int {
	teamIds := []int{}
	seenTeams := make(map[int]bool)

	for _, roster := range seriesRosters {

		team, ok := roster["team"].(map[string]interface{})
		if ok {
			teamId := int(team["id"].(float64))
			_, teamAdded := seenTeams[teamId]
			if !teamAdded {
				seenTeams[teamId] = true
				teamIds = append(teamIds, teamId)
			}
		}
	}
	return teamIds
}
