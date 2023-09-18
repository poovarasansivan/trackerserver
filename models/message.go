package models

type MessageModel struct {
    Id       int    `json:"id"`
    Content  string `json:"content"`
    Response string `json:"response"`
}
