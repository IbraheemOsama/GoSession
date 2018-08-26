package models

// Movie : the properties in mongodb document
type Movie struct {
	Name string `json:"name"`
	Year string `json:"year"`
}
