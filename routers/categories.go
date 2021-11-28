package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/go-solid-principle/controllers"
	"github.com/wsaefulloh/go-solid-principle/repos"
	"github.com/wsaefulloh/go-solid-principle/services"
)

func CategoryRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewCategory(db)
	service := services.CategoryServices(repo)
	cr := controllers.NewCategory(service)
	route := r.PathPrefix("/category").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.Add).Methods(http.MethodPost)
	route.HandleFunc("", cr.Add).Methods(http.MethodPost)
	route.HandleFunc("/{id}", cr.Delete).Methods(http.MethodDelete)
	route.HandleFunc("/{id}", cr.Update).Methods(http.MethodPut)
}
