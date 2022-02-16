package models

type Recipe struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	Ingredients     []Ingredient `json:"ingredients"`
	Steps           []string     `json:"steps"`
	PreparationTime int          `json:"preparation_time"`
	CookingTime     int          `json:"cooking_time"`
}
