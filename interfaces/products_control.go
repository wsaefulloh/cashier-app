package interfaces

import "net/http"

type ProductControl interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	SearchbyName(w http.ResponseWriter, r *http.Request)
	SearchbyCategory(w http.ResponseWriter, r *http.Request)
	GetbyCategory(w http.ResponseWriter, r *http.Request)
	GetbyDateASC(w http.ResponseWriter, r *http.Request)
	GetbyDateDESC(w http.ResponseWriter, r *http.Request)
	GetbyPriceDESC(w http.ResponseWriter, r *http.Request)
	GetbyPriceASC(w http.ResponseWriter, r *http.Request)
}
