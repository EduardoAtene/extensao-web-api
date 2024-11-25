package main

import (
	"log"
	"net/http"
	"your-project/internal/config"
	"your-project/internal/handlers"
	"your-project/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	db := config.SetupDB()
	r := mux.NewRouter()

	// Handlers
	authHandler := handlers.NewAuthHandler(db)

	// Rotas
	r.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Middleware
	handler := middleware.CORS(r)

	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
