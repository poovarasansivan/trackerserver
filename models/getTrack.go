package models

type TrackModal struct {
	Id        int     `json:"id"`
	Car       int     `json:"car"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Course    float32 `json:"course"`
}
