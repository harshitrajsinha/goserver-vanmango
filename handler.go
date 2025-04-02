package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/harshitrajsinha/goserver-vanmango/test"
	_ "github.com/lib/pq"
)

// Function to load data to database via schema file
// func loadDataToDatabase(dbClient *sql.DB, filename string) error {

// 	// Read file content
// 	sqlFile, err := os.ReadFile(filename)
// 	if err != nil {
// 		return err
// 	}

// 	// Execute file content (queries)
// 	_, err = dbClient.Exec(string(sqlFile))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/", test.HandleHomeRoute).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on PORT ", port)
	router.ServeHTTP(w, r)
}
