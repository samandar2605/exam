package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"exam/models"
	"exam/storage/repo"

	_ "github.com/lib/pq"
)

type GetCarsParams struct {
	Limit  int
	Page   int
	Search string
}

type GetCarsResponse struct {
	Cars  []models.Cars
	Count int
}

type dbCars struct {
	db *sql.DB
}

func NewCars(connStr *sql.DB) repo.RepoCars {
	return dbCars{db: connStr}
}

func (b dbCars) CreateCar(auto models.Cars) (int, error) {
	var carId int
	tx, err := b.db.Begin()
	if err != nil {
		log.Fatalf("Error at creating transaction %v", err)
	}

	carQuery := `
		INSERT INTO cars(
			image_url,
			marka,
			model,
			color,
			mileage_km,
			made_year,
			cost
		) values($1,$2,$3,$4,$5,$6,$7)
		RETURNING id
	`
	var result *sql.Rows
	result, err = tx.Query(
		carQuery,
		auto.ImageUrl,
		auto.Marka,
		auto.Model,
		auto.Color,
		auto.MileageKm,
		auto.MadeYear,
		auto.Cost,
	)

	if err != nil {
		tx.Rollback()
		log.Fatalf("Error at Insert method\n%v", err)
	}

	if err = result.Scan(&carId); err != nil {
		log.Fatalf("Error at Scanning element \n%v", err)
	}

	imageQuery := `
		INSERT INTO images(
			cars_id,
			image_url,
			squence_number
		) values($1,$2,$3)
	`

	for _, image := range auto.Images {
		_, err = tx.Exec(
			imageQuery,
			image.CarsId,
			image.ImageUrl,
			image.SequenceNumber,
		)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Error at Insert images\n%v", err)
		}
	}

	tx.Commit()
	return carId, nil
}

func (b dbCars) ReadCar(id int) (models.Cars, error) {
	var car models.Cars
	car.Images = make([]*models.Images, 0)

	query := `
		SELECT 
			image_url,
			marka,
			model,
			color,
			mileage_km,
			made_year,
			cost,
		from cars
		where id=$1
	`

	result := b.db.QueryRow(query, id)
	err := result.Scan(
		&car.ImageUrl,
		&car.Marka,
		&car.Model,
		&car.Color,
		&car.MileageKm,
		&car.MadeYear,
		&car.Cost,
	)

	if err != nil {
		log.Fatalf("error at scanning %v", err)
		return models.Cars{}, err
	}

	imageQuery := `
		SELECT
			id,
			image_url,
			sequence_number
		FROM  images
		WHERE cars_id=$1
	`
	rows, err := b.db.Query(
		imageQuery,
		id)
	if err != nil {
		log.Fatalf("Error at query images\n%v", err)
		return models.Cars{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var image models.Images
		err := rows.Scan(
			&image.ID,
			&image.ImageUrl,
			&image.SequenceNumber,
		)
		if err != nil {
			log.Fatalf("Error at scan image\n%v", err)
			return models.Cars{}, err
		}
		car.Images = append(car.Images, &image)
	}

	return car, nil
}

func (b dbCars) ReadCarAll(limit, page int, search string) ([]models.Cars, int, error) {
	var cars []models.Cars

	filter := ""
	if search != "" {
		filter = fmt.Sprintf("WHERE model ilike '%s'", "%"+search+"%")
	}

	query := `
		SELECT
			id,
			image_url,
			marka,
			model,
			color,
			mileage_km,
			made_year,
			cost
		from cars
	` + filter + `
		order by made_year desc
		limit $1 offset $2
	`
	offset := (page - 1) * limit
	rows, err := b.db.Query(query, limit, offset)
	if err != nil {
		log.Fatalf("error at select filter car\n%v", err)
		return []models.Cars{}, 0, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var car models.Cars
		err := rows.Scan(
			&car.ID,
			&car.ImageUrl,
			&car.Marka,
			&car.Model,
			&car.Color,
			&car.MileageKm,
			&car.MadeYear,
			&car.Cost,
		)
		count++
		if err != nil {
			log.Fatalf("Error at scaning car elements\n%v", err)
			return []models.Cars{}, 0, err
		}
		cars = append(cars, car)
	}
	return cars, count, nil
}

func (b dbCars) UpdateCar(crd models.Cars) error {

	tx, err := b.db.Begin()
	if err != nil {
		log.Fatalf("Error at creating transaction\n%v", err)
	}

	query := `
		UPDATE cars SET
			image_url=$1,
			marka=$2,
			model=$3,
			color=$4,
			mileage_km=$5,
			made_year=$6,
			cost=$7
		where id=&8
	`
	result, err := b.db.Exec(
		query,
		crd.ImageUrl,
		crd.Marka,
		crd.Model,
		crd.Color,
		crd.MileageKm,
		crd.MadeYear,
		crd.Cost,
	)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Error at update method\n%v", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	queryDeleteImage := `DELETE FROM images where cars_id=$1`
	_, err = b.db.Exec(queryDeleteImage, crd.ID)
	if err != nil {
		return err
	}

	imageQuery := `
		INSERT INTO images(
			cars_id,
			image_url,
			squence_number
		) values($1,$2,$3)
	`
	for _, image := range crd.Images {
		_, err = tx.Exec(
			imageQuery,
			image.CarsId,
			image.ImageUrl,
			image.SequenceNumber,
		)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Error at Insert images\n%v", err)
		}
	}
	return nil
}

func (b dbCars) DeleteCar(id int) error {
	tx, err := b.db.Begin()
	if err != nil {
		log.Fatalf("Error at creating transaction\n%v", err)
	}

	queryDeleteImage := `DELETE FROM images where cars_id=$1`
	_, err = tx.Exec(queryDeleteImage, id)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Error at deleting element\n%v", err)
	}

	queryDeleteCar := `DELETE FROM cars WHERE id=$1`
	result, err := tx.Exec(queryDeleteCar, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
