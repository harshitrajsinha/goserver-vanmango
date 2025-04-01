package handler

import (
	"encoding/json"
	"strings"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func VerifyEngineRequestBody(body []byte, checkType int) bool {

	var doesKeyDisplacementExists bool
	var doesKeyCylinderExists bool
	var doesKeyMaterialExists bool
	var data map[string]interface{}

	_ = json.Unmarshal([]byte(body), &data)
	if _, displacementExists := data["displacement"]; !displacementExists {
		doesKeyDisplacementExists = false
	} else {
		doesKeyDisplacementExists = true
	}
	if _, cylinderExists := data["no-of-cylinders"]; !cylinderExists {
		doesKeyCylinderExists = false
	} else {
		doesKeyCylinderExists = true
	}
	if _, materialExists := data["material"]; !materialExists {
		doesKeyMaterialExists = false
	} else {
		doesKeyMaterialExists = true
	}

	if checkType == 1 {
		return doesKeyCylinderExists && doesKeyDisplacementExists && doesKeyMaterialExists
	} else if checkType == 0 {
		return doesKeyCylinderExists || doesKeyDisplacementExists || doesKeyMaterialExists
	} else {
		return false
	}
}

func VerifyVanRequestBody(body []byte, checkType int) bool {

	var doesKeyNameExists bool
	var doesKeyBrandExists bool
	var doesKeyDescriptionExists bool
	var doesKeyCategoryExists bool
	var doesKeyFuelTypeExists bool
	var doesKeyEngineIDExists bool
	var doesKeyPriceExists bool
	var doesKeyImageURLExists bool
	var data map[string]interface{}

	_ = json.Unmarshal(body, &data)

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

	if checkType == 1 {
		return doesKeyNameExists && doesKeyBrandExists && doesKeyDescriptionExists && doesKeyCategoryExists && doesKeyFuelTypeExists && doesKeyEngineIDExists && doesKeyPriceExists && doesKeyImageURLExists
	} else if checkType == 0 {
		return doesKeyNameExists || doesKeyBrandExists || doesKeyDescriptionExists || doesKeyCategoryExists || doesKeyFuelTypeExists || doesKeyEngineIDExists || doesKeyPriceExists || doesKeyImageURLExists
	} else {
		return false
	}
}

// Use when there is another dedicated version, otherwise server will automatically handle via 404 not found
func CheckAPIVersion(urlPath string) bool {
	segment := strings.Split(urlPath, "/")
	apiVersion := segment[2]
	if apiVersion != "v1" {
		return false
	} else {
		return true
	}
}
