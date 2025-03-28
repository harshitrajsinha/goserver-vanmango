package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/harshitrajsinha/van-man-go/models"
	"github.com/harshitrajsinha/van-man-go/service"
)

type VanHandler struct {
	service service.VanServiceInterface
}

func NewVanHandler(service service.VanServiceInterface) *VanHandler {
	return &VanHandler{
		service: service,
	}
}

// Function to get Engine data by ID
func (v *VanHandler) GetVanByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get van id
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if id is valid uuid
	result, _ := uuid.Parse(id)
	if result.Version() != 4 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "Invalid engine ID"})
		log.Println("Invalid Engine ID")
		return
	}

	// Get data from service layer
	resp, err := v.service.GetVanById(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	// Send response
	var respData []interface{}
	respData = append(respData, resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{Code: http.StatusOK, Data: respData})
	log.Println("Data populated successfully")
}

// Function to get All Engine list
func (v *VanHandler) GetAllVan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get data from service layer
	resp, err := v.service.GetAllVan(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	// Send response
	var respData []interface{}
	respData = append(respData, resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{Code: http.StatusOK, Data: respData})
	log.Println("Data populated successfully")
}

// Function to create Engine
func (v *VanHandler) CreateVan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	var vanReq models.Van
	err = json.Unmarshal(body, &vanReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	// Verify request body
	if err := models.ValidateVaneReq(vanReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: fmt.Sprintf(err.Error())})
		log.Println(err)
		return
	}

	// Pass data to service layer to create engine
	createdVan, err := v.service.CreateVan(ctx, &vanReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Code: http.StatusCreated, Message: createdVan["message"]})
	log.Println("Data inserted successfully")
}

// Function to update Engine data by id
func (v *VanHandler) UpdateVan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get van id
	params := mux.Vars(r)
	id := params["id"]

	// Check if id is valid uuid
	result, _ := uuid.Parse(id)
	if result.Version() != 4 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "Invalid van ID"})
		log.Println("Invalid van ID")
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	var vanReq models.Van
	err = json.Unmarshal(body, &vanReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	// Pass data to service layer to update van
	updatedVan, err := v.service.UpdateVan(ctx, id, &vanReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	if updatedVan > 0 {
		// data is updated successfully
		log.Println("Data updated successfully!")
		// Get the updated result
		v.GetVanByID(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "No data present for provided Van ID"})
		log.Println("value of updatedVan is ", updatedVan)
		return
	}
}

// Function to delete Van
func (v *VanHandler) DeleteVan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get van id
	params := mux.Vars(r)
	id := params["id"]

	// Check if id is valid uuid
	result, _ := uuid.Parse(id)
	if result.Version() != 4 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "Invalid van ID"})
		log.Println("Invalid van ID")
		return
	}

	deletedVan, err := v.service.DeleteVan(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while deleting data"})
		log.Println(err)
		return
	}

	if deletedVan > 0 {
		// data is deleted successfully
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
		log.Println("value of updatedVan is ", deletedVan)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "No data present for provided Van ID"})
		log.Println("value of updatedVan is ", deletedVan)
		return
	}

}
