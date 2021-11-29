package repos

import (
	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type HistoryRepMock struct {
	Mocks mock.Mock
}

var dataHistory = models.History{
	Id:      1,
	Invoice: "a",
	Cashier: "a",
	Date:    "Now",
	Order:   "1",
	Count:   2,
	Total:   2,
}

var dataHistories = models.Histories{
	dataHistory,
}

func (pr *HistoryRepMock) FindAll() (*models.Histories, error) {
	args := pr.Mocks.Called()

	if args == nil {
		return nil, nil
	}

	return &dataHistories, nil
}

func (pr *HistoryRepMock) Save(prod *models.History) error {
	args := pr.Mocks.Called(prod)

	if args == nil {
		return nil
	}

	return nil
}

func (pr *HistoryRepMock) Remove(id string) error {
	args := pr.Mocks.Called(id)

	if args == nil {
		return nil
	}

	return nil
}

func (pr *HistoryRepMock) Edit(prod *models.History, id string) error {
	args := pr.Mocks.Called(prod, id)

	if args == nil {
		return nil
	}

	return nil
}
