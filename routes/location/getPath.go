package location

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"track/config"
	"track/functions"
	"track/models"
)

func GetPath(w http.ResponseWriter, r *http.Request) {
	functions.AllowHeaders(w)

	var myLocations []models.LocationModal
	var temp models.LocationModal
	var input models.PathInput

	err := json.NewDecoder(r.Body).Decode(&input)

	var response map[string]interface{}
	if err != nil {
		response = map[string]interface{}{
			"success": false,
			"error":   "Invalid Request",
		}
		functions.Response(w, response)
		return
	}
	row, err := config.Database.Query("SELECT latitude,longitude FROM location_paths WHERE (from_location =? AND to_location=?) or (from_location =? AND to_location=?) ORDER BY sort_order ASC", input.FromLocation, input.ToLocation, input.ToLocation, input.FromLocation)
	if err != nil {
		if err == sql.ErrNoRows {
			response = map[string]interface{}{
				"success": false,
				"error":   "No Records",
			}
		} else {
			response = map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			}
		}
	} else {

		for row.Next() {
			err := row.Scan(&temp.Latitude, &temp.Longitude)
			if err != nil {
				panic(err.Error())
			}

			tempRow := models.LocationModal{
				Latitude:  temp.Latitude,
				Longitude: temp.Longitude,
			}
			myLocations = append(myLocations, tempRow)

		}

		response = map[string]interface{}{
			"success": true,
			"data":    myLocations,
		}

	}
	functions.Response(w, response)
}
