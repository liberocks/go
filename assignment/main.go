package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"

	"github.com/liberocks/go/assignment/handler"
	"github.com/liberocks/go/assignment/utils"
)

func main() {
	dbUrl := os.Getenv("DB_URL")

	// For the sake of simplicity, we will run the migration up every time the server starts.
	// Ideally, we should only run the migration up once manually.
	utils.Up(dbUrl)

	r := mux.NewRouter()

	r.HandleFunc("/", handler.RootHandler)
	r.HandleFunc("/orders", handler.OrderHandlers)
	r.HandleFunc("/orders/{id}", handler.OrderDetailHandlers)

	http.Handle("/", r)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
