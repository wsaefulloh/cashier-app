package repos

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type CategoryRepMock struct {
	Mocks mock.Mock
}

var dataCategory = models.Category{
	Name:      "testing",
	Id:        1,
	CreatedAt: time.Now(),
	UpdateAt:  time.Now(),
}

var dataCategories = models.Categories{
	dataCategory,
}

func (pr *CategoryRepMock) FindAll() (*models.Categories, error) {
	args := pr.Mocks.Called()

	if args == nil {
		return nil, nil
	}

	return &dataCategories, nil
}

func (pr *CategoryRepMock) Save(prod *models.Category) error {
	args := pr.Mocks.Called(prod)

	if args == nil {
		return nil
	}

	return nil
}

func (pr *CategoryRepMock) Remove(id string) error {
	args := pr.Mocks.Called(id)

	if args == nil {
		return nil
	}

	return nil
}

func (pr *CategoryRepMock) Edit(prod *models.Category, id string) error {
	args := pr.Mocks.Called(prod, id)

	if args == nil {
		return nil
	}

	return nil
}
