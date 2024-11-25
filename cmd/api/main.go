package main

import (
	"log"
	"net/http"

	"github.com/EduardoAtene/extensao-web-api/internal/config"
	handlers "github.com/EduardoAtene/extensao-web-api/internal/handler"
	middleware "github.com/EduardoAtene/extensao-web-api/internal/middlaware"
	"github.com/gorilla/mux"
)

func main() {
	db := config.SetupDB()
	r := mux.NewRouter()

	authHandler := handlers.NewAuthHandler(db)
	discardHandler := handlers.NewDiscardHandler(db)

	// Rotas
	r.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Rotas que precisam de login
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/discard", discardHandler.CreateDiscard).Methods("POST")
	protected.HandleFunc("/rewards", discardHandler.GetUserRewards).Methods("GET")

	// // Middleware
	handler := middleware.CORS(r)

	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
