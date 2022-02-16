package models

type Ingredient struct {
	Name        string  `json:"name"`
	Amount      float32 `json:"amount"`
	Measurement string  `json:"measurement"`
}
