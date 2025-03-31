package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/harshitrajsinha/goserver-vanmango/models"
	"github.com/harshitrajsinha/goserver-vanmango/service"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		service: service,
	}
}

// Function to get Engine data by ID
func (e *EngineHandler) GetEngineByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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
	resp, err := e.service.GetEngineByID(ctx, id)
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
func (e *EngineHandler) GetAllEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get data from service layer
	resp, err := e.service.GetAllEngine(ctx)
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
func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
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

	var engineReq models.Engine
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	// Verify request body
	if err := models.ValidateEngineReq(engineReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: err.Error()})
		log.Println(err)
		return
	}

	// Pass data to service layer to create engine
	createdEngine, err := e.service.CreateEngine(ctx, &engineReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Code: http.StatusCreated, Message: createdEngine["message"]})
	log.Println("Data inserted successfully")
}

// Function to update Engine data by id
func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	// Check if id is valid uuid
	result, _ := uuid.Parse(id)
	if result.Version() != 4 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "Invalid engine ID"})
		log.Println("Invalid Engine ID")
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

	var engineReq models.Engine
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	// Pass data to service layer to update engine
	updatedEngine, err := e.service.UpdateEngine(ctx, id, &engineReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}

	if updatedEngine > 0 {
		// data is updated successfully
		log.Println("Data updated successfully!")
		// Get the updated result
		e.GetEngineByID(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "No data present for provided Engine ID"})
		log.Println("value of updatedEngine is ", updatedEngine)
		return
	}
}

// Function to delete Engine
func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	// Check if id is valid uuid
	result, _ := uuid.Parse(id)
	if result.Version() != 4 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "Invalid engine ID"})
		log.Println("Invalid Engine ID")
		return
	}

	deletedEngine, err := e.service.DeleteEngine(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while deleting data"})
		log.Println(err)
		return
	}

	if deletedEngine > 0 {
		// data is deleted successfully
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
		log.Println("value of updatedEngine is ", deletedEngine)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "No data present for provided Engine ID"})
		log.Println("value of updatedEngine is ", deletedEngine)
		return
	}

}
