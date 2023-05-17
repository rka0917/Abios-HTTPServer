package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rka0917/Abios-HTTPServer/models"
)

// Handler used for handling the /series endpoint
func LiveSeriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	liveSeries, err := models.GetLiveSeries()

	if err != nil {
		http.Error(w, "Error occurred while retrieving data.", http.StatusInternalServerError)
	}

	jsonData, err := json.MarshalIndent(liveSeries, "", "  ")

	if err != nil {
		http.Error(w, "Error occurred while retrieving data.", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonData))
}
