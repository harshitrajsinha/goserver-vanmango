package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/harshitrajsinha/van-man-go/models"
	"github.com/harshitrajsinha/van-man-go/service"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

type Response struct{
	Code int `json:"code"`
	Message string `json:"message,omitempty"`
	Data []interface{} `json:"data,omitempty"`
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler{
	return &EngineHandler{
		service: service,
	}
}

func (e *EngineHandler) GetEngineByID(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if id is valid uuid
	result, _ := uuid.Parse(id)
	if result.Version() != 4{
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusBadRequest, Message: "Invalid engine ID"})
		log.Println("Invalid Engine ID")
		return
	}

	// Get data from service layer
	resp, err := e.service.GetEngineByID(ctx, id)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println(err)
		return
	}
	_, err = json.Marshal(resp)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Code: http.StatusInternalServerError, Message: "Error occured while reading data"})
		log.Println("handler/GetEngineByID : ", err)
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

func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil{
		log.Println("Error reading request: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.Engine
	err = json.Unmarshal(body, &engineReq)
	if err != nil{
		log.Println("Error during unmarshal: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdEngine, err := e.service.CreateEngine(ctx, &engineReq)
	if err != nil{
		log.Println("Error while creating Engine: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, err := json.Marshal(createdEngine)
	if err != nil{
		log.Println("Error during unmarshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(resBody)
}

func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	
	body, err := io.ReadAll(r.Body)
	if err != nil{
		log.Println("Error reading request: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.Engine
	err = json.Unmarshal(body, &engineReq)
	if err != nil{
		log.Println("Error during unmarshal: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedEngine, err := e.service.UpdateEngine(ctx, id, &engineReq)
	if err != nil{
		log.Println("Error while updating Engine: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, err := json.Marshal(updatedEngine)
	if err != nil{
		log.Println("Error during unmarshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resBody)
	return

}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]
	
	// body, err := io.ReadAll(r.Body)
	// if err != nil{
	// 	log.Println("Error reading request: ", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// var engineReq models.Engine
	// err = json.Unmarshal(body, &engineReq)
	// if err != nil{
	// 	log.Println("Error during unmarshal: ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	deletedEngine, err := e.service.DeleteEngine(ctx, id)
	if err != nil{
		log.Println("Error while updating Engine: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// check if engine deleted successfully
	// if deletedEngine.engineID == uuid.Nil{
	// 	w.WriteHeader(http.StatusNotFound)
	// 	response := map[string]string{"error": "Engine not found"}
	// 	jsonResponse, _ := json.Marshal(response)
	// 	_, _ = w.Write(jsonResponse)
	// 	return
	// }

	resBody, err := json.Marshal(deletedEngine)
	if err != nil{
		log.Println("Error during unmarshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resBody)

}
