package main

import (
	"encoding/json"
	"maglev.onebusaway.org/internal/models"
	"net/http"
	"time"
)

// Declare a handler which writes a JSON response with information about the
// current time.
func (app *application) currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	if app.requestHasInvalidAPIKey(r) {
		app.invalidAPIKeyResponse(w, r)
		return
	}

	timeData := models.NewCurrentTimeData(time.Now())
	response := models.NewResponse(http.StatusOK, timeData, "OK")

	// Set content type and write response
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
