package main

import (
	"net/http"
	"github.com/PhreakZ3r0/TODO-user-service/internal/handler"
	"github.com/PhreakZ3r0/TODO-user-service/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	userHandler := handler.NewUserHandler()

	mux.HandleFunc("/register", middleware.AuthMiddleware(userHandler.Register))
	mux.HandleFunc("/profile", middleware.AuthMiddleware(userHandler.Profile))

	http.ListenAndServe(":8080", mux)
}
