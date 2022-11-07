package main

import (
	"database/sql"
	"exam/config"
	"exam/models"
	"exam/storage/postgres"
	"fmt"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("postgres", config.ConnStr())

	if err != nil {
		log.Fatalf("failed open database: %v", err)
	}

	database := postgres.NewCars(db)

	id, err := database.CreateCar(models.Cars{
		ImageUrl:  "dang",
		Marka:     "Chevrolet",
		Model:     "Malibu",
		Color:     "Qora",
		MileageKm: 234567,
		MadeYear:  time.Time{2009-11-10 23:00:00 +0000 UTC m=+0.000000001},
		Cost:      54654654,
		Images: []*models.Images{
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

	fmt.Println(id)

	if err != nil {
		log.Fatalf("failed to created car %v", err)
	}

	product, err := database.ReadCar(1)
	fmt.Println(product)

	result,n, err := database.ReadCarAll(10, 1, "Malibu")

	fmt.Print(result,n)

	err=database.UpdateCar(models.Cars{
		ImageUrl:  "dang",
		Marka:     "Chevrolet",
		Model:     "Malibu",
		Color:     "Qora",
		MileageKm: 234567,
		MadeYear:  time.Time{2009-11-10 23:00:00 +0000 UTC m=+0.000000001},
		Cost:      54654654,
		Images: []*models.Images{
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

	if err!=nil{
		log.Fatalf("Update'da xato %v",err)
	}
	
	err=database.DeleteCar(1)
	if err!=nil{
		log.Fatalf("Deleteda xatolik!")
	}

}
