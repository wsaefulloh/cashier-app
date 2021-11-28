package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/interfaces"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type users struct {
	rp interfaces.UserServices
}

func New(rps interfaces.UserServices) *users {
	return &users{rps}
}

func (us *users) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	data, err := us.rp.FindAll()

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	data.Send(w)
}

func (us *users) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	validate := validator.New()
	var body models.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	err = validate.Struct(body)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		helpers.Respone(w, validationErrors.Error(), 400, true)
		return
	}

	user, err := us.rp.GetUsername(body.UserName)

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	if len([]rune(user)) != 0 {
		helpers.Respone(w, "Username sudah dipakai", 400, true)
		return
	}

	data := models.CreateUser()
	data.Name = body.Name
	data.UserName = body.UserName
	data.Email = body.Email
	data.Role = body.Role
	newPass, errPass := helpers.HashPassword(body.Password)
	if errPass != nil {
		// fmt.Println("gagal hashing")
		helpers.Respone(w, err.Error(), 500, true)
		return
	}
	data.Password = newPass

	result, err := us.rp.AddData(data)

	if err != nil {
		helpers.Respone(w, err.Error(), 500, true)
		return
	}

	result.Send(w)
}
