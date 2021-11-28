package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/go-solid-principle/controllers"
	"github.com/wsaefulloh/go-solid-principle/repos"
	"github.com/wsaefulloh/go-solid-principle/services"
)

func AuthRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewUsers(db)
	service := services.UserServices(repo)
	cr := controllers.Auth(service)

	route := r.PathPrefix("/auth").Subrouter()
	route.HandleFunc("/", cr.Login).Methods(http.MethodPost)
	route.HandleFunc("/chek", cr.OpenToken).Methods(http.MethodGet)
}
