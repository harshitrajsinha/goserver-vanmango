package models

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type Van struct {
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	FuelType    string    `json:"fuel-type"`
	EngineID    uuid.UUID `json:"engine-id"`
	Price       float64   `json:"price"`
}

// Validate name
func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

// Validate brand name
func validateBrandName(brand string) error {
	if brand == "" {
		return errors.New("brand name is required")
	}
	return nil
}

// Validate fuel type
func validateFuelType(fuel string) error {

	for _, value := range [3]string{"petrol", "diesel", "gasoline"} {
		if strings.ToLower(fuel) == value {
			return nil
		}
	}
	return errors.New("fuel type must be one of following - ['petrol', 'diesel', 'gasoline']")
}

func validateEngine(engine Engine) error {
	err := ValidateEngineReq(engine)
	return err
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}

func ValidateVaneReq(vanRequest Van) error {
	var err error

	if err = validateName(vanRequest.Name); err != nil {
		return err
	}

	if err = validateBrandName(vanRequest.Brand); err != nil {
		return err
	}

	if err = validateFuelType(vanRequest.FuelType); err != nil {
		return err
	}

	// if err = validateEngine(vanRequest.Engine); err != nil{
	// 	return err
	// }

	if err = validatePrice(vanRequest.Price); err != nil {
		return err
	}

	return nil
}
