package models

type Images struct {
	ID             int    `json:"id" db:"id"`
	CarsId         int    `json:"cars_id" db:"cars_id"`
	ImageUrl       string `json:"image_url" db:"image_url"`
	SequenceNumber int    `json:"sequence_number" db:"sequence_number"`
}
