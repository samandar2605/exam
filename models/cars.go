package models

import "time"

type Cars struct {
	ID        int       `json:"id" db:"id"`
	ImageUrl  string    `json:"image_url" db:"image_url"`
	Marka     string    `json:"marka" db:"marka"`
	Model     string    `json:"model" db:"model"`
	Color     string    `json:"color" db:"color"`
	MileageKm int       `json:"mileage_km" db:"mileage_km"`
	MadeYear  time.Time `json:"made_year" db:"made_year"`
	Cost      float64   `json:"cost" db:"cost"`
	Images    []Images
}
