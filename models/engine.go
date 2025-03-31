package models

import (
	"encoding/json"
	"errors"
	"strings"
)

type Engine struct {
	Displacement  int64  `json:"displacement"`
	NoOfCylinders int    `json:"no-of-cylinders"`
	Material      string `json:"material"`
}

// 1500cc – 4000cc
func validateDisplacement(displacement int64) error {
	if displacement < 1500 || displacement > 4000 {
		return errors.New("displacement must fall within the range of 1500-4000")
	}
	return nil
}

// [4,6,8]
func validateCylinderNo(cylinders int) error {
	for _, value := range [3]int{4, 6, 8} {
		if cylinders == value {
			return nil
		}
	}
	return errors.New("no. of cylinders must be one of following - [4, 6, 8]")
}

// aluminium || iron
func validateMaterial(material string) error {
	if strings.ToLower(material) == "aluminium" || strings.ToLower(material) == "iron" {
		return nil
	}
	return errors.New("material must be one of following - ['aluminium', 'iron']")
}

// Function to check which key exists in request body
func verifyRequestBody(request []byte) [3]bool {

	var doesKeyDisplacementExists bool
	var doesKeyCylinderExists bool
	var doesKeyMaterialExists bool
	var data map[string]interface{}

	_ = json.Unmarshal([]byte(request), &data)

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

	return [3]bool{doesKeyDisplacementExists, doesKeyCylinderExists, doesKeyMaterialExists}
}

func ValidateEngineCrReq(engineRequest Engine) error {
	var err error

	if err = validateDisplacement(engineRequest.Displacement); err != nil {
		return err
	}

	if err = validateCylinderNo(engineRequest.NoOfCylinders); err != nil {
		return err
	}

	if err = validateMaterial(engineRequest.Material); err != nil {
		return err
	}

	return nil
}

func ValidateEngineUpReq(request []byte) error {
	var err error
	var engineRequest Engine

	_ = json.Unmarshal(request, &engineRequest)

	// Check which key exists and verify accordingly
	doesKeyExists := verifyRequestBody(request)

	if doesKeyExists[0] { // if "displacement" exists in request body
		if err = validateDisplacement(engineRequest.Displacement); err != nil {
			return err
		}
	}
	if doesKeyExists[1] { // if "no-of-cylinders" exists in request body
		if err = validateCylinderNo(engineRequest.NoOfCylinders); err != nil {
			return err
		}
	}
	if doesKeyExists[2] { // if "material" exists in request body
		if err = validateMaterial(engineRequest.Material); err != nil {
			return err
		}
	}

	return nil
}
