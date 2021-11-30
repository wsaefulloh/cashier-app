package controllers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/models"
	"github.com/wsaefulloh/go-solid-principle/services"
)

var productsService = services.ProductServiceMock{Mocks: mock.Mock{}}
var prodsControl = NewProduct(&productsService)

func TestGetAllProduct(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/", prodsControl.GetAll)
	productsService.Mocks.On("FindAll")

	r.ServeHTTP(w, httptest.NewRequest("GET", "/test/", nil))

	// var prods models.Products
	var respon helpers.Response

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("konversi ke struck gagal")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.Equal(t, 200, respon.Status, "Status code must be 200")
	assert.False(t, respon.IsError, "isError must be false")

}

func TestSaveProduct(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/add/", prodsControl.Add)

	Newprod := &models.Product{
		Name:     "testing",
		Price:    2000,
		Category: "1",
	}

	productsService.Mocks.On("AddData", Newprod)

	byteData, _ := json.Marshal(Newprod)

	req := httptest.NewRequest("POST", "/test/add/", bytes.NewBuffer(byteData))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	var respon helpers.Response

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("konversi ke struck gagal")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.Equal(t, 200, respon.Status, "Status code must be 200")
	assert.False(t, respon.IsError, "isError must be false")

}

func TestUpdateProduct(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/update/{id}", prodsControl.Update)

	Newprod := &models.Product{
		Name:     "testing",
		Price:    2000,
		Category: "1",
	}

	productsService.Mocks.On("UpdateData", Newprod, "0")

	byteData, _ := json.Marshal(Newprod)

	req := httptest.NewRequest("PUT", "/test/update/0", bytes.NewBuffer(byteData))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	var respon helpers.Response

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("konversi ke struck gagal")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.Equal(t, 200, respon.Status, "Status code must be 200")
	assert.False(t, respon.IsError, "isError must be false")

}

func TestDeleteProduct(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/test/del/{id}", prodsControl.Delete)

	productsService.Mocks.On("Remove", "0")

	req := httptest.NewRequest("DELETE", "/test/del/0", nil)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	var respon helpers.Response

	err := json.Unmarshal(w.Body.Bytes(), &respon)

	if err != nil {
		t.Fatal("konversi ke struck gagal")
	}

	assert.Equal(t, 200, w.Code, "Status code must be 200")
	assert.Equal(t, 200, respon.Status, "Status code must be 200")
	assert.False(t, respon.IsError, "isError must be false")

}
