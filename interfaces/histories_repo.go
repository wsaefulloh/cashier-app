package interfaces

import "github.com/wsaefulloh/go-solid-principle/models"

type RepoHistory interface {
	FindAll() (*models.Histories, error)
	Save(histo *models.History) error
	Remove(id string) error
	Edit(histo *models.History, id string) error
}
