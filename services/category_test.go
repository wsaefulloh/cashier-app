package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/models"
	"github.com/wsaefulloh/go-solid-principle/repos"
)

var categoryRepo = repos.CategoryRepMock{Mocks: mock.Mock{}}
var cateService = category_service{rep: &categoryRepo}

func TestCategory(t *testing.T) {
	t.Run("Get All Category", func(t *testing.T) {
		categoryRepo.Mocks.On("FindAll")
		data, err := cateService.FindAll()
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Save Category", func(t *testing.T) {
		prod := models.Category{
			Name:      "testing",
			Id:        1,
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
		}
		categoryRepo.Mocks.On("Save", &prod)
		data, err := cateService.AddData(&prod)
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Update Category", func(t *testing.T) {
		prod := models.Category{
			Name:      "testing",
			Id:        1,
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
		}
		categoryRepo.Mocks.On("Edit", &prod, "1")
		data, err := cateService.UpdateData(&prod, "1")
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Delete Category", func(t *testing.T) {
		categoryRepo.Mocks.On("Remove", "1")
		data, err := cateService.Remove("1")
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})
}
