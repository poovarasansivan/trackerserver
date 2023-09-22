package track

import (
	"encoding/json"
	"net/http"

)

type Message struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Response string `json:"response"`
}

var messages []Message

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var newMessage Message
	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMessage.ID = len(messages) + 1
	messages = append(messages, newMessage)

	w.WriteHeader(http.StatusCreated)
}

