package services

import (
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/interfaces"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type user_service struct {
	rep interfaces.RepoUser
}

func UserServices(reps interfaces.RepoUser) *user_service {
	return &user_service{reps}
}

func (pr *user_service) FindAll() (*helpers.Response, error) {
	data, err := pr.rep.FindAll()
	if err != nil {
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}

func (pr *user_service) AddData(prod *models.User) (*helpers.Response, error) {
	error := pr.rep.Save(prod)

	if error != nil {
		return nil, error
	}

	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  "Product berhasil ditambahkan",
	}

	return &respon, nil
}

func (pr *user_service) GetPass(username string) (string, error) {
	data, err := pr.rep.GetPass(username)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (pr *user_service) GetEmail(username string) (string, error) {
	data, err := pr.rep.GetEmail(username)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (pr *user_service) GetUsername(username string) (string, error) {
	data, err := pr.rep.GetUsername(username)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (pr *user_service) GetRole(username string) (string, error) {
	data, err := pr.rep.GetRole(username)
	if err != nil {
		return "", err
	}
	return data, nil
}
