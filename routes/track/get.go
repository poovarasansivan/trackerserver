package track

import (
	"database/sql"
	"net/http"
	"track/config"
	"track/functions"
	"track/models"
)

func GetTrack(w http.ResponseWriter, r *http.Request) {
	functions.AllowHeaders(w)

	var myLocations []models.TrackModal
	var temp models.TrackModal

	var response map[string]interface{}
	row, err := config.Database.Query("select id,car_id,latitude,longitude,course from location order by id desc limit 1")
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
			err := row.Scan(&temp.Id, &temp.Car, &temp.Latitude, &temp.Longitude, &temp.Course)
			if err != nil {
				panic(err.Error())
			}

			tempRow := models.TrackModal{
				Id:        temp.Id,
				Car:       temp.Car,
				Latitude:  temp.Latitude,
				Longitude: temp.Longitude,
				Course:    temp.Course,
			}
			myLocations = append(myLocations, tempRow)

		}

		response = map[string]interface{}{
			"success": true,
			"data":    temp,
		}

	}

	functions.Response(w, response)
}
