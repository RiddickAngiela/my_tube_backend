package main

import (
    "log"
    "github.com/joho/godotenv"
    "net/http"
    "github.com/gorilla/mux"
    "my_tube_backend/handlers"
    "my_tube_backend/middleware"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

    protected := r.PathPrefix("/admin").Subrouter()
    protected.Use(middleware.AuthMiddleware)
    // Add your admin routes here

    http.ListenAndServe(":8080", r)
}
