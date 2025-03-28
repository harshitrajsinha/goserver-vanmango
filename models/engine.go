package models

import (
	"errors"
	"strings"
)

type Engine struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylinders int `json:"no-of-cylinders"`
	Material      string `json:"material"`
}

// 1500cc – 4000cc
func validateDisplacement(displacement int64) error {
	if displacement < 1500 || displacement > 4000 {
		return errors.New("displacement must fall within the range of 1500cc-4000cc")
	}
	return nil
}

// [4,6,8]
func validateCylinderNo(cylinders int) error {
	for _, value := range [3]int{4, 6, 8}{
		if cylinders == value{
			return nil
		}
	}
	return errors.New("possible no. of cylinders - [4, 6, 8]")
}

// aluminium || iron
func validateMaterial(material string) error {
	if strings.ToLower(material) == "aluminium"  || strings.ToLower(material) == "iron" {
		return nil
	}
	return errors.New("material could be either aluminium or iron")
}

func ValidateEngineReq(engineRequest Engine) error{
	var err error
	
	if err = validateDisplacement(engineRequest.Displacement); err != nil{
		return err
	}

	if err = validateCylinderNo(engineRequest.NoOfCylinders); err != nil{
		return err
	}

	if err = validateMaterial(engineRequest.Material); err != nil{
		return err
	}

	return nil
}