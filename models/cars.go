package models

import "time"

type Cars struct {
	ID        int
	ImageUrl  string
	Marka     string
	Model     string
	Color     string
	MileageKm int
	MadeYear  time.Time
	Cost      float64
	Images    []Images
}
