package logic

import (
	"go_mongo_test/internal/database"
	"go_mongo_test/internal/domain"
)

func GetCars() ([]domain.CarInternal, error) {
	return database.FindCars()
}

func PutCar(c domain.CarInternal) error {
	return database.AddCar(c)
}
