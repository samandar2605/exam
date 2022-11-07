package main

import (
	"database/sql"
	"exam/config"
	"exam/models"
	"exam/storage/postgres"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", config.ConnStr())

	if err != nil {
		log.Fatalf("failed open database: %v", err)
	}

	database := postgres.NewCars(db)

	id, err := database.CreateCar(models.Cars{
		ImageUrl:  "test",
		Marka:     "Chevrolet",
		Model:     "Malibu",
		Color:     "Qora",
		MileageKm: 234567,
		MadeYear:  time.Now(),
		Cost:      54654654,
		Images: []models.Images{
			{
				CarsId:         1,
				ImageUrl:       "test",
				SequenceNumber: 1,
			},
			{
				CarsId:         11111,
				ImageUrl:       "test1",
				SequenceNumber: 2,
			},
		},
	})

	fmt.Println(id)

	if err != nil {
		log.Fatalf("failed to created car %v", err)
	}

	product, err := database.ReadCar(1)
	fmt.Println(product)

	result, n, err := database.ReadCarAll(1, 1, "k5")

	fmt.Print(result, n)

	err = database.UpdateCar(models.Cars{
		ID:        2,
		ImageUrl:  "dang",
		Marka:     "Chevrolet",
		Model:     "Malibuuuu",
		Color:     "Qora",
		MileageKm: 234567,
		MadeYear:  time.Now(),
		Cost:      54654654,
		Images: []models.Images{
			{
				CarsId:         1,
				ImageUrl:       "dang",
				SequenceNumber: 1,
			},
			{
				CarsId:         2,
				ImageUrl:       "bang",
				SequenceNumber: 3,
			},
		},
	})

	err = database.DeleteCar(1)
	if err != nil {
		log.Fatalf("Deleteda xatolik!")
	}
	fmt.Println(err)
}
