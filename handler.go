package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/gorilla/mux"
	"github.com/harshitrajsinha/goserver-vanmango/driver"
	apiV1 "github.com/harshitrajsinha/goserver-vanmango/handler/v1"
	"github.com/harshitrajsinha/goserver-vanmango/service"
	"github.com/harshitrajsinha/goserver-vanmango/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Function to load data to database via schema file
func loadDataToDatabase(dbClient *sql.DB, filename string) error {

	// Read file content
	sqlFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Execute file content (queries)
	_, err = dbClient.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Server is functioning"})
}

func Handler(w http.ResponseWriter, r *http.Request) {

	var err error
	var sqlSchemaFile string = "store/schema.sql"
	var dbClient *sql.DB

	// panic recovery
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error occured: ", r)
			debug.PrintStack()
		}
	}()

	// Load env file data to sys env (development)
	_ = godotenv.Load()

	// initialize database connection
	var message string
	err = driver.InitDB()
	if err != nil {
		panic(err)
	} else {
		log.Println("Successfully connected to database")
	}
	log.Println(message)

	// close db connection
	defer func() {
		err = driver.CloseDB()
		if err != nil {
			panic(err)
		}
	}()

	// Get instance of database client
	dbClient = driver.GetDB()

	// Load data into database
	err = loadDataToDatabase(dbClient, sqlSchemaFile)
	if err != nil {
		panic(err)
	} else {
		log.Println("SQL file executed successfully!")
	}
	router := mux.NewRouter()
	router.HandleFunc("/", handleHomeRoute).Methods("GET")

	// Initialize engine constructors
	engineStore := store.NewEngineStore(dbClient)
	engineService := service.NewEngineService(engineStore)
	_ = apiV1.NewEngineHandler(engineService)

	// Initialize van constructors
	vanStore := store.NewVanStore(dbClient)
	vanService := service.NewVanService(vanStore)
	_ = apiV1.NewVanHandler(vanService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on PORT ", port)
	router.ServeHTTP(w, r)
}
