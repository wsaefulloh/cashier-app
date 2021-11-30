package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wsaefulloh/go-solid-principle/configs/db"
)

func New() *mux.Router {
	r := mux.NewRouter()

	//inisialisasi endpoint
	r.HandleFunc("/", simpleHandlers).Methods(http.MethodGet)
	mainRutes := r.PathPrefix("/api/v1").Subrouter().StrictSlash(false)
	// mainRutes.HandleFunc("/foo", fooHandlers).Methods(http.MethodGet)

	//inisialisasi dbms
	dbms, _ := db.New()
	clientRedis := db.Client()
	db.Connect(clientRedis)

	UserRoute(mainRutes, dbms)
	AuthRoute(mainRutes, dbms)
	ProductRoute(mainRutes, dbms)
	CategoryRoute(mainRutes, dbms)
	HistoryRoute(mainRutes, dbms)

	mainRutes.Use(mux.CORSMethodMiddleware(mainRutes))

	return mainRutes
}

func simpleHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}
