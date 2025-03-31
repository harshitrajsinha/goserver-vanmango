package handler

import "encoding/json"

type Response struct {
	Code    int           `json:"code"`
	Message string        `json:"message,omitempty"`
	Data    []interface{} `json:"data,omitempty"`
}

func VerifyRequestBody(body []byte, checkType int) bool {

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
