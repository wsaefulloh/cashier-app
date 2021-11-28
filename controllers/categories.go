package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/interfaces"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type categories struct {
	rp interfaces.CategoryServices
}

func NewCategory(rps interfaces.CategoryServices) *categories {
	return &categories{rps}
}

func (cate *categories) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := cate.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	data.Send(w)
}

func (cate *categories) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.Category
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data := models.CreateCategory()
	data.Name = body.Name

	result, err := cate.rp.AddData(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil disimpan"))
	result.Send(w)
}

func (cate *categories) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	data, err := cate.rp.Remove(vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil dihapus"))
	data.Send(w)
}

func (cate *categories) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var body models.Category
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	vars := mux.Vars(r)

	data := models.CreateCategory()
	data.Name = body.Name

	result, err := cate.rp.UpdateData(data, vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// w.Write([]byte("Data berhasil diedit"))
	result.Send(w)
}
