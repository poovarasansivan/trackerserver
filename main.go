package main

import (
    "log"
    "net/http"
    "track/config"
    "track/functions"
    "track/routes/location"
    "track/routes/track"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    config.ConnectDB()
    defer config.Database.Close()

    http.HandleFunc("/", functions.CheckAuth)
    http.HandleFunc("/getTrack", track.GetTrack)
    http.HandleFunc("/getLocation", location.GetLocation)
    http.HandleFunc("/getPath", location.GetPath)

    // Pass db to GetMessages
	http.HandleFunc("/getMessages", func(w http.ResponseWriter, r *http.Request) {
		track.GetMessages(w, r, config.Database)
	})
	

    // Start the server
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
