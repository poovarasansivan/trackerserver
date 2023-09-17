package location

import (
	"database/sql"
	"net/http"
	"track/config"
	"track/functions"
	"track/models"
)

func GetLocation(w http.ResponseWriter, r *http.Request) {
	functions.AllowHeaders(w)

	var myLocations []models.LocationModal
	var temp models.LocationModal

	var response map[string]interface{}
	row, err := config.Database.Query("select * from m_location where status ='1'")
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
			err := row.Scan(&temp.Id, &temp.Name, &temp.Latitude, &temp.Longitude, &temp.Status)
			if err != nil {
				panic(err.Error())
			}

			tempRow := models.LocationModal{
				Id:        temp.Id,
				Name:      temp.Name,
				Latitude:  temp.Latitude,
				Longitude: temp.Longitude,
				Status:    temp.Status,
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
