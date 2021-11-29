package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/models"
	"github.com/wsaefulloh/go-solid-principle/repos"
)

var historyRepo = repos.HistoryRepMock{Mocks: mock.Mock{}}
var histoService = history_service{rep: &historyRepo}

func TestHistory(t *testing.T) {
	t.Run("Get All History", func(t *testing.T) {
		historyRepo.Mocks.On("FindAll")
		data, err := histoService.FindAll()
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Save History", func(t *testing.T) {
		prod := models.History{
			Id:      1,
			Invoice: "a",
			Cashier: "a",
			Date:    "Now",
			Order:   "1",
			Count:   2,
			Total:   2,
		}
		historyRepo.Mocks.On("Save", &prod)
		data, err := histoService.AddData(&prod)
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Update History", func(t *testing.T) {
		prod := models.History{
			Id:      1,
			Invoice: "a",
			Cashier: "a",
			Date:    "Now",
			Order:   "1",
			Count:   2,
			Total:   2,
		}
		historyRepo.Mocks.On("Edit", &prod, "1")
		data, err := histoService.UpdateData(&prod, "1")
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Delete History", func(t *testing.T) {
		historyRepo.Mocks.On("Remove", "1")
		data, err := histoService.Remove("1")
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})
}
