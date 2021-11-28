package interfaces

import (
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type HistoryServices interface {
	FindAll() (*helpers.Response, error)
	AddData(prod *models.History) (*helpers.Response, error)
	Remove(id string) (*helpers.Response, error)
	UpdateData(prod *models.History, id string) (*helpers.Response, error)
}
