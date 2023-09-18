package track

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "log"
)

type Message struct {
    ID       int    `json:"id"`
    Content  string `json:"content"`
    Response string `json:"response"`
}

func getResponse(db *sql.DB, input string) (string, error) {
	log.Println( input)
    row := db.QueryRow("SELECT response FROM messages WHERE content = ?", input)
    var response string
    err := row.Scan(&response)
    if err != nil {
        return "", err
    }
    return response, nil
}

func GetMessages(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    input := r.FormValue("content")
    botMessage, err := getResponse(db, input)

    if err != nil {
        log.Println("Error getting response:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if botMessage == "" {
        botMessage = "Sorry, I don't understand."
    }

    var messages []Message
    messages = append(messages, Message{Content: input, Response: "user"})
    messages = append(messages, Message{Content: botMessage, Response: "Driver"})

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(messages)
}
