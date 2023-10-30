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
	switch r.Method {
	case "POST":
		h.RegisterUser(r,w)
	case "GET":
		h.GetUser(r,w)
	case "PUT":
		h.UpdateUser(r,w)
	case "DELETE":
		h.DeleteUser(r,w)
	default:
		http.Error(w, "Method Call Not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	
}

func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	// Handle fetching user profile
}
