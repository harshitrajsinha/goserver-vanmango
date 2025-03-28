package store

import (
	"database/sql"
)

type Store struct{
	db *sql.DB
}

func new(db *sql.DB) Store{
	return Store{db: db}
}

// func (s Store) GetCarById(ctx context.Context, id string)(models.Car, error){
	
// }

// func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool)(models.Car, error){
	
// }

// func (s Store) CreateCar(ctx context.Context, carReq *models.CarRequest)(models.Car, error){
	
// }

// func (s Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest)(models.Car, error){
	
// }

// func (s Store) DeleteCar(ctx context.Context, id string)(models.Car, error){
	
// }