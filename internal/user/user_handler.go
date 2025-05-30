package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", h.GetUsers).Methods("GET")
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.FindAll()
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	user := h.service.Create(input.Name)
	json.NewEncoder(w).Encode(user)
}