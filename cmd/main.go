package main

import (
	"database/sql"
	v1 "exam/api/v1"
	"exam/config"
	"exam/models"
	"exam/storage/postgres"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	router.POST("/car",)
	
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
	router.Run()

	// err = database.UpdateCar(models.Cars{
	// 	ID:        2,
	// 	ImageUrl:  "dang",
	// 	Marka:     "Chevrolet",
	// 	Model:     "Malibuuuu",
	// 	Color:     "Qora",
	// 	MileageKm: 234567,
	// 	MadeYear:  time.Now(),
	// 	Cost:      54654654,
	// 	Images: []models.Images{
	// 		{
	// 			CarsId:         1,
	// 			ImageUrl:       "dang",
	// 			SequenceNumber: 1,
	// 		},
	// 		{
	// 			CarsId:         2,
	// 			ImageUrl:       "bang",
	// 			SequenceNumber: 3,
	// 		},
	// 	},
	// })

	// err = database.DeleteCar(1)
	// if err != nil {
	// 	log.Fatalf("Deleteda xatolik!")
	// }
	// fmt.Println(err)
}

// func main() {
	// router.GET("/test", getTests)
// 	router.GET("/test/:id", GetTest)
	// router.POST("/test", addTest)
// 	router.Run("localhost:8080")
// }
