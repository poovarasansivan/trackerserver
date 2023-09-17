package models

type LocationModal struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Status    string  `json:"status"`
}
