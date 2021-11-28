package interfaces

import "net/http"

type ProductControl interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}
