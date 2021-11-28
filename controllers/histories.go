package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/interfaces"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type histories struct {
	rp interfaces.HistoryServices
}

func NewHistory(rps interfaces.HistoryServices) *histories {
	return &histories{rps}
}

func (histo *histories) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := histo.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	data.Send(w)
}

func (histo *histories) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.History
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data := models.CreateHistory()
	data.Invoice = body.Invoice
	data.Cashier = body.Cashier
	data.Order = body.Order
	data.Count = body.Count
	data.Total = 0

	result, err := histo.rp.AddData(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil disimpan"))
	result.Send(w)
}

func (histo *histories) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	data, err := histo.rp.Remove(vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil dihapus"))
	data.Send(w)
}

func (histo *histories) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.History
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	vars := mux.Vars(r)

	data := models.CreateHistory()
	data.Invoice = body.Invoice
	data.Cashier = body.Cashier
	data.Order = body.Order
	data.Count = body.Count
	data.Total = 0

	result, err := histo.rp.UpdateData(data, vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil diedit"))
	result.Send(w)
}
