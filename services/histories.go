package services

import (
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/interfaces"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type history_service struct {
	rep interfaces.RepoHistory
}

func HistoryServices(reps interfaces.RepoHistory) *history_service {
	return &history_service{reps}
}

func (pr *history_service) FindAll() (*helpers.Response, error) {
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

func (pr *history_service) AddData(histo *models.History) (*helpers.Response, error) {
	error := pr.rep.Save(histo)

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

func (pr *history_service) UpdateData(histo *models.History, id string) (*helpers.Response, error) {
	error := pr.rep.Edit(histo, id)

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

func (pr *history_service) Remove(id string) (*helpers.Response, error) {
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
