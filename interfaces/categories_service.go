package interfaces

import (
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type CategoryServices interface {
	FindAll() (*helpers.Response, error)
	AddData(prod *models.Category) (*helpers.Response, error)
	Remove(id string) (*helpers.Response, error)
	UpdateData(prod *models.Category, id string) (*helpers.Response, error)
}
