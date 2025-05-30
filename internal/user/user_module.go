package user

import "github.com/gorilla/mux"

func RegisterUserModule(r *mux.Router) {
	repo := NewInMemoryUserRepo()
	service := NewUserService(repo)
	handler := NewUserHandler(service)
	handler.RegisterRoutes(r)
}