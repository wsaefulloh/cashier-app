package interfaces

import (
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type UserServices interface {
	FindAll() (*helpers.Response, error)
	AddData(prod *models.User) (*helpers.Response, error)
	GetPass(username string) (string, error)
	GetEmail(username string) (string, error)
	GetUsername(username string) (string, error)
	GetRole(username string) (string, error)
}
