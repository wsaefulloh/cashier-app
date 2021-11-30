package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/wsaefulloh/go-solid-principle/configs/db"
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/interfaces"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type products struct {
	rp interfaces.ProductServices
}

func NewProduct(rps interfaces.ProductServices) *products {
	return &products{rps}
}

func (pro *products) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		fmt.Println(err)
		return
	}

	byteData, _ := json.Marshal(data.Result)

	cacheErr := db.Client().Set(db.Ctx, "product", byteData, 30*time.Second).Err()
	if cacheErr != nil {
		fmt.Println(cacheErr)
		return
	}

	json.Unmarshal(byteData, &data.Result)

	// helpers.Respone(w, data, 200, false)

	data.Send(w)
}

func (pro *products) GetAllNoRedis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.FindAll()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		fmt.Println(err)
		return
	}

	// helpers.Respone(w, data, 200, false)

	data.Send(w)
}

func (pro *products) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var body models.Product
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data := models.CreateProduct()
	data.Name = body.Name
	data.Price = body.Price
	data.Category = body.Category

	result, err := pro.rp.AddData(data)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	cacheErr := db.Client().Del(db.Ctx, "product").Err()
	if cacheErr != nil {
		fmt.Println(cacheErr)
		return
	}

	// w.Write([]byte("Data berhasil disimpan"))
	result.Send(w)
}

func (pro *products) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	data, err := pro.rp.Remove(vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	cacheErr := db.Client().Del(db.Ctx, "product").Err()
	if cacheErr != nil {
		fmt.Println(cacheErr)
		return
	}

	// w.Write([]byte("Data berhasil dihapus"))
	data.Send(w)
}

func (pro *products) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body models.Product
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	vars := mux.Vars(r)
	data := models.CreateProduct()
	data.Name = body.Name
	data.Price = body.Price
	data.Category = body.Category

	result, err := pro.rp.UpdateData(data, vars["id"])

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	cacheErr := db.Client().Del(db.Ctx, "product").Err()
	if cacheErr != nil {
		fmt.Println(cacheErr)
		return
	}

	// w.Write([]byte("Data berhasil diedit"))
	result.Send(w)
}

func (pro *products) SearchbyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	name_prod := strings.Join(vars["name"], " ")
	data, err := pro.rp.SearchbyName(name_prod)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// fmt.Println(name_prod)
	// json.NewEncoder(w).Encode(&data)
	data.Send(w)
}

func (pro *products) SearchbyCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	name_category := strings.Join(vars["category"], " ")
	data, err := pro.rp.SearchbyCategory(name_category)

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// fmt.Println(name_category)
	// json.NewEncoder(w).Encode(&data)
	data.Send(w)
}

func (pro *products) GetbyCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.SortbyCategory()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data.Send(w)
}

func (pro *products) GetbyDateASC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.SortbyDateASC()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	data.Send(w)
}

func (pro *products) GetbyDateDESC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.SortbyDateDESC()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	data.Send(w)
}

func (pro *products) GetbyPriceDESC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.SortbyPriceDESC()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	data.Send(w)
}

func (pro *products) GetbyPriceASC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data, err := pro.rp.SortbyPriceASC()

	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	// json.NewEncoder(w).Encode(&data)
	data.Send(w)
}
