package v1

import (
	"errors"
	"exam/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCarAll(context *gin.Context, cars []models.Cars) {
	context.IndentedJSON(http.StatusOK, cars)
}

func AddCar(context *gin.Context)models.Cars {
	var NewCar models.Cars
	if err := context.BindJSON(&NewCar); err != nil {
		log.Fatalf("Error at AddCar api'da\n%v", err)
	}
	context.IndentedJSON(http.StatusOK, NewCar)
	return NewCar
}

func GetCar(id int, cars []models.Cars) (*models.Cars, error) {
	for _, t := range cars {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("quiz no found")
}

func GetCarById(context *gin.Context, cars []models.Cars) {
	id := context.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf(err.Error())
	}
	test, err := GetCar(ID, cars)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "test no found"})
		return
	}

	context.IndentedJSON(http.StatusOK, test)
}
