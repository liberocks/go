package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/liberocks/go/assignment/handler"
	"github.com/liberocks/go/assignment/middleware"
	"github.com/liberocks/go/assignment/utils"
)

func main() {
	dbUrl := os.Getenv("DB_URL")

	// For the sake of simplicity, we will run the migration up every time the server starts.
	// Ideally, we should only run the migration up once manually.
	utils.Up(dbUrl)

	r := mux.NewRouter()

	r.HandleFunc("/", handler.RootHandler)
	r.HandleFunc("/sign-in", handler.SignInHandlers)
	r.HandleFunc("/sign-up", handler.SignUpHandlers)
	r.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(handler.OrderHandlers)))
	r.Handle("/orders/{id}", middleware.AuthMiddleware(http.HandlerFunc(handler.OrderDetailHandlers)))

	http.Handle("/", r)

	log.Info().Msg("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
