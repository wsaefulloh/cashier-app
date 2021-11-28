package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/go-solid-principle/controllers"
	"github.com/wsaefulloh/go-solid-principle/middleware"
	"github.com/wsaefulloh/go-solid-principle/repos"
	"github.com/wsaefulloh/go-solid-principle/services"
)

func HistoryRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewHistory(db)
	service := services.HistoryServices(repo)
	cr := controllers.NewHistory(service)
	route := r.PathPrefix("/history").Subrouter()
	route.HandleFunc("/", middleware.Validate(cr.GetAll)).Methods(http.MethodGet)
	route.HandleFunc("", middleware.Validate(cr.GetAll)).Methods(http.MethodGet)
	route.HandleFunc("/", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("/{id}", middleware.Validate(cr.Delete)).Methods(http.MethodDelete)
	route.HandleFunc("/{id}", middleware.Validate(cr.Update)).Methods(http.MethodPut)
}
