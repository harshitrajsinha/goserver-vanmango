package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/harshitrajsinha/van-man-go/driver"
	engineHandler "github.com/harshitrajsinha/van-man-go/handler"
	engineService "github.com/harshitrajsinha/van-man-go/service"
	engineStore "github.com/harshitrajsinha/van-man-go/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func executeSchemaFile(db *sql.DB, filename string) error{
	sqlFile, err := os.ReadFile(filename)
	if err != nil{
		return err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil{
		return err
	}
	fmt.Println("SQL file executed successfully!")
	return nil
}

func main(){
	err := godotenv.Load()

	if err != nil{
		log.Fatal("Error loading env file")
	}

	driver.InitDB()
	defer driver.CloseDB()
	
	db := driver.GetDB()
	
	engineStore := engineStore.NewEngineStore(db)
	engineService := engineService.NewEngineService(engineStore)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	router := mux.NewRouter()
	
	schemaFile := "./store/schema.sql"
	if err := executeSchemaFile(db, schemaFile); err != nil{
		log.Fatal("Error while executing schema file: ", err)
	}
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Server is functioning"})
	}).Methods("GET")
	router.HandleFunc("/engine/{id}", engineHandler.GetEngineByID).Methods("GET")
	router.HandleFunc("/engine", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	log.Println("Server listening on PORT ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}