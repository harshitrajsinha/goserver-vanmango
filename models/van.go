package models

import (
	"encoding/json"
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
	Price       int64     `json:"price"`
	ImageURL    string    `json:"image-url"`
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

// Validate description
func validateDescription(description string) error {
	if description == "" {
		return errors.New("description is required")
	}
	return nil
}

// Validate description
func validateCategory(category string) error {
	for _, value := range [3]string{"simple", "rugged", "luxury"} {
		if strings.ToLower(category) == value {
			return nil
		}
	}
	return errors.New("category must be one of following - ['simple', 'rugged', 'luxury']")
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

// Validate engine ID
func validateEngineID(engineID uuid.UUID) error {
	if engineID.Version() != 4 {
		return errors.New("engineID is required")
	}
	return nil
}

// validate van price
func validatePrice(price int64) error {
	if price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}

// Validate image URL
func validateImageURL(image string) error {
	if image == "" {
		return errors.New("image is required")
	}
	return nil
}

// Function to check which key exists in request body
func verifyVanRequestKeys(request []byte) [8]bool {

	var doesKeyNameExists bool
	var doesKeyBrandExists bool
	var doesKeyDescriptionExists bool
	var doesKeyCategoryExists bool
	var doesKeyFuelTypeExists bool
	var doesKeyEngineIDExists bool
	var doesKeyPriceExists bool
	var doesKeyImageURLExists bool
	var data map[string]interface{}

	_ = json.Unmarshal([]byte(request), &data)

	if _, nameExists := data["name"]; !nameExists {
		doesKeyNameExists = false
	} else {
		doesKeyNameExists = true
	}
	if _, brandExists := data["brand"]; !brandExists {
		doesKeyBrandExists = false
	} else {
		doesKeyBrandExists = true
	}
	if _, descriptionExists := data["description"]; !descriptionExists {
		doesKeyDescriptionExists = false
	} else {
		doesKeyDescriptionExists = true
	}
	if _, categoryExists := data["category"]; !categoryExists {
		doesKeyCategoryExists = false
	} else {
		doesKeyCategoryExists = true
	}
	if _, fuelTypeExists := data["fuel-type"]; !fuelTypeExists {
		doesKeyFuelTypeExists = false
	} else {
		doesKeyFuelTypeExists = true
	}
	if _, engineIdExists := data["engine-id"]; !engineIdExists {
		doesKeyEngineIDExists = false
	} else {
		doesKeyEngineIDExists = true
	}
	if _, priceExists := data["price"]; !priceExists {
		doesKeyPriceExists = false
	} else {
		doesKeyPriceExists = true
	}
	if _, imageUrlExists := data["image-url"]; !imageUrlExists {
		doesKeyImageURLExists = false
	} else {
		doesKeyImageURLExists = true
	}

	return [8]bool{doesKeyNameExists, doesKeyBrandExists, doesKeyDescriptionExists, doesKeyCategoryExists, doesKeyFuelTypeExists, doesKeyEngineIDExists, doesKeyPriceExists, doesKeyImageURLExists}
}

func ValidateVanReq(vanRequest Van) error {
	var err error

	if err = validateName(vanRequest.Name); err != nil {
		return err
	}

	if err = validateBrandName(vanRequest.Brand); err != nil {
		return err
	}

	if err = validateDescription(vanRequest.Description); err != nil {
		return err
	}

	if err = validateCategory(vanRequest.Category); err != nil {
		return err
	}

	if err = validateFuelType(vanRequest.FuelType); err != nil {
		return err
	}

	if err = validateEngineID(vanRequest.EngineID); err != nil {
		return err
	}

	if err = validatePrice(vanRequest.Price); err != nil {
		return err
	}

	if err = validateImageURL(vanRequest.ImageURL); err != nil {
		return err
	}

	return nil
}

func ValidateVanPatchReq(request []byte) error {
	var err error
	var vanRequest Van

	_ = json.Unmarshal(request, &vanRequest)

	// Check which key exists and verify accordingly
	doesKeyExists := verifyVanRequestKeys(request)

	if doesKeyExists[0] { // if "name" exists in request body
		if err = validateName(vanRequest.Name); err != nil {
			return err
		}
	}
	if doesKeyExists[1] { // if "brand" exists in request body
		if err = validateBrandName(vanRequest.Brand); err != nil {
			return err
		}
	}
	if doesKeyExists[2] { // if "description" exists in request body
		if err = validateDescription(vanRequest.Description); err != nil {
			return err
		}
	}
	if doesKeyExists[3] { // if "category" exists in request body
		if err = validateCategory(vanRequest.Category); err != nil {
			return err
		}
	}
	if doesKeyExists[4] { // if "fuel-type" exists in request body
		if err = validateFuelType(vanRequest.FuelType); err != nil {
			return err
		}
	}
	if doesKeyExists[5] { // if "engine-id" exists in request body
		if err = validateEngineID(vanRequest.EngineID); err != nil {
			return err
		}
	}
	if doesKeyExists[6] { // if "price" exists in request body
		if err = validatePrice(vanRequest.Price); err != nil {
			return err
		}
	}
	if doesKeyExists[7] { // if "image-url" exists in request body
		if err = validateImageURL(vanRequest.ImageURL); err != nil {
			return err
		}
	}

	return nil
}
