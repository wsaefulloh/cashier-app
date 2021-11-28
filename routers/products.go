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

func ProductRoute(r *mux.Router, db *sql.DB) {
	repo := repos.NewProduct(db)
	service := services.ProductServices(repo)
	cr := controllers.NewProduct(service)
	route := r.PathPrefix("/products").Subrouter()
	route.HandleFunc("/", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("", cr.GetAll).Methods(http.MethodGet)
	route.HandleFunc("/date/desc", cr.GetbyDateDESC).Methods(http.MethodGet)
	route.HandleFunc("/date/asc", cr.GetbyDateASC).Methods(http.MethodGet)
	route.HandleFunc("/price/desc", cr.GetbyPriceDESC).Methods(http.MethodGet)
	route.HandleFunc("/price/asc", cr.GetbyPriceASC).Methods(http.MethodGet)
	route.HandleFunc("/category", cr.GetbyCategory).Methods(http.MethodGet)
	route.HandleFunc("/search/name", cr.SearchbyName).Methods(http.MethodGet)
	route.HandleFunc("/search/category", cr.SearchbyCategory).Methods(http.MethodGet)
	route.HandleFunc("/", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("", middleware.Validate(cr.Add)).Methods(http.MethodPost)
	route.HandleFunc("/{id}", middleware.Validate(cr.Delete)).Methods(http.MethodDelete)
	route.HandleFunc("/{id}", middleware.Validate(cr.Update)).Methods(http.MethodPut)
}
