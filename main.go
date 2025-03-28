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
	"github.com/harshitrajsinha/van-man-go/handler"
	"github.com/harshitrajsinha/van-man-go/middleware"
	"github.com/harshitrajsinha/van-man-go/service"
	"github.com/harshitrajsinha/van-man-go/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func executeSchemaFile(db *sql.DB, filename string) error {
	sqlFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	fmt.Println("SQL file executed successfully!")
	return nil
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Server is functioning"})
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()

	// Load data to table
	schemaFile := "./store/schema.sql"
	if err := executeSchemaFile(db, schemaFile); err != nil {
		log.Fatal("Error while executing schema file: ", err)
	}

	// Initialize engine constructors
	engineStore := store.NewEngineStore(db)
	engineService := service.NewEngineService(engineStore)
	engineHandler := handler.NewEngineHandler(engineService)

	// Initialize van constructors
	vanStore := store.NewVanStore(db)
	vanService := service.NewVanService(vanStore)
	vanHandler := handler.NewVanHandler(vanService)

	router := mux.NewRouter()
	router.HandleFunc("/", handleHomeRoute).Methods("GET")
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")

	protectedRouter := router.PathPrefix("/").Subrouter()
	protectedRouter.Use(middleware.AuthMiddleware)

	// Routes for Engine
	protectedRouter.HandleFunc("/engine/{id}", engineHandler.GetEngineByID).Methods("GET")
	protectedRouter.HandleFunc("/engine", engineHandler.GetAllEngine).Methods("GET")
	protectedRouter.HandleFunc("/engine", engineHandler.CreateEngine).Methods("POST")
	protectedRouter.HandleFunc("/engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	protectedRouter.HandleFunc("/engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	// Routes for Van
	protectedRouter.HandleFunc("/van/{id}", vanHandler.GetVanByID).Methods("GET")
	protectedRouter.HandleFunc("/van", vanHandler.GetAllVan).Methods("GET")
	protectedRouter.HandleFunc("/van", vanHandler.CreateVan).Methods("POST")
	protectedRouter.HandleFunc("/van/{id}", vanHandler.UpdateVan).Methods("PUT")
	protectedRouter.HandleFunc("/van/{id}", vanHandler.DeleteVan).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on PORT ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
