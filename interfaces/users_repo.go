package interfaces

import "github.com/wsaefulloh/go-solid-principle/models"

type RepoUser interface {
	FindAll() (*models.Users, error)
	Save(user *models.User) error
	GetPass(username string) (string, error)
	GetEmail(username string) (string, error)
	GetUsername(username string) (string, error)
	GetRole(username string) (string, error)
}
