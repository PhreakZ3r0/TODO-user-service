package handler

import (
	"net/http"
)

type UserHandler struct {
	// Dependency Injection of services
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Handle user registration
}

func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	// Handle fetching user profile
}
