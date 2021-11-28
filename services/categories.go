package services

import (
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/interfaces"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type category_service struct {
	rep interfaces.RepoCategory
}

func CategoryServices(reps interfaces.RepoCategory) *category_service {
	return &category_service{reps}
}

func (pr *category_service) FindAll() (*helpers.Response, error) {
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

func (pr *category_service) AddData(cate *models.Category) (*helpers.Response, error) {
	error := pr.rep.Save(cate)

	if error != nil {
		return nil, error
	}

	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  "Category berhasil ditambahkan",
	}

	return &respon, nil
}

func (pr *category_service) UpdateData(cate *models.Category, id string) (*helpers.Response, error) {
	error := pr.rep.Edit(cate, id)

	if error != nil {
		return nil, error
	}

	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  "Update category berhasil",
	}

	return &respon, nil
}

func (pr *category_service) Remove(id string) (*helpers.Response, error) {
	error := pr.rep.Remove(id)

	if error != nil {
		return nil, error
	}

	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  "Category berhasil dihapus",
	}

	return &respon, nil
}
