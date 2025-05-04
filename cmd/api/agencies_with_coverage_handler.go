package main

import (
	"maglev.onebusaway.org/internal/models"
	"net/http"
)

func (app *application) agenciesWithCoverageHandler(w http.ResponseWriter, r *http.Request) {
	// help fill me in.
	// Create the agency data
	agency := map[string]interface{}{
		"agencyId": "unitrans",
		"lat":      38.555308499999995,
		"latSpan":  0.03564300000000031,
		"lon":      -121.73599,
		"lonSpan":  0.10499999999998977,
	}

	// Create the agency reference
	agencyRef := map[string]interface{}{
		"disclaimer":     "",
		"email":          "",
		"fareUrl":        "",
		"id":             "unitrans",
		"lang":           "en",
		"name":           "Unitrans",
		"phone":          "530-752-BUSS",
		"privateService": false,
		"timezone":       "America/Los_Angeles",
		"url":            "http://unitrans.ucdavis.edu",
	}

	// Create references with the agency
	references := models.ReferencesModel{
		Agencies:   []interface{}{agencyRef},
		Routes:     []interface{}{},
		Situations: []interface{}{},
		StopTimes:  []interface{}{},
		Stops:      []interface{}{},
		Trips:      []interface{}{},
	}

	// Create the data structure
	data := map[string]interface{}{
		"limitExceeded": false,
		"list":          []interface{}{agency},
		"references":    references,
	}

	// Create the response
	response := models.NewOKResponse(data)

	app.sendResponse(w, r, response)
}
