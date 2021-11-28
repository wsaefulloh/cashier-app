package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/go-solid-principle/controllers"
	"github.com/wsaefulloh/go-solid-principle/repos"
	"github.com/wsaefulloh/go-solid-principle/services"
)

func UserRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewUsers(db)
	service := services.UserServices(repo)
	cr := controllers.New(service)
	route := r.PathPrefix("/users").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/", cr.Add).Methods(http.MethodPost)
}
