package main

import (
	"maglev.onebusaway.org/internal/models"
	"net/http"
)

func (app *application) agenciesWithCoverageHandler(w http.ResponseWriter, r *http.Request) {
	agencies := app.gtfsManager.GetAgencies()

	agenciesWithCoverage := make([]models.AgencyCoverage, len(agencies))
	agencyReferences := make([]models.AgencyReference, len(agencies))

	for _, a := range agencies {
		agenciesWithCoverage = append(
			agenciesWithCoverage,
			models.NewAgencyCoverage(a.Id, 0.0, 0.0, 0.0, 0.0),
		)

		agencyReferences = append(
			agencyReferences,
			models.NewAgencyReference(
				a.Id,
				a.Name,
				a.Url,
				a.Timezone,
				a.Language,
				a.Phone,
				a.Email,
				"",
				"",
				false,
			),
		)
	}

	// Create references with the agency
	references := models.ReferencesModel{
		Agencies:   agencyReferences,
		Routes:     []interface{}{},
		Situations: []interface{}{},
		StopTimes:  []interface{}{},
		Stops:      []interface{}{},
		Trips:      []interface{}{},
	}

	// Create the data structure
	data := map[string]interface{}{
		"limitExceeded": false,
		"list":          agenciesWithCoverage,
		"references":    references,
	}

	// Create the response
	response := models.NewOKResponse(data)

	app.sendResponse(w, r, response)
}
