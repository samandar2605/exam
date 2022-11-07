package repo

import (
	"exam/models"
)

type RepoCars interface {
	CreateCar(auto models.Cars) (int, error)
	ReadCar(id int) (models.Cars, error)
	ReadCarAll(limit int, page int, search string) ([]models.Cars,int, error)
	UpdateCar(crd models.Cars)error
	DeleteCar(id int) error
}
