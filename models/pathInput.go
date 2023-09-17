package models

type PathInput struct{
	FromLocation int `json:"from_location"`
	ToLocation int `json:"to_location"`
}