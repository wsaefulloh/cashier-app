package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/models"
	"github.com/wsaefulloh/go-solid-principle/repos"
)

var productsRepo = repos.ProductRepMock{Mocks: mock.Mock{}}
var prodsService = product_service{rep: &productsRepo}

func TestProduct(t *testing.T) {
	t.Run("Get All Product", func(t *testing.T) {
		productsRepo.Mocks.On("FindAll")
		data, err := prodsService.FindAll()
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Save Product", func(t *testing.T) {
		prod := models.Product{
			Name:      "testing",
			Price:     2000,
			Id:        1,
			Category:  "1",
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
		}
		productsRepo.Mocks.On("Save", &prod)
		data, err := prodsService.AddData(&prod)
		res := data.Result
		assert.Equal(t, "Product berhasil ditambahkan", res, "Result must be 'Product berhasil ditambahkan'")
		assert.Nil(t, err)
	})

	t.Run("Update Product", func(t *testing.T) {
		prod := models.Product{
			Name:      "testing",
			Price:     2000,
			Id:        1,
			Category:  "1",
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
		}
		productsRepo.Mocks.On("Edit", &prod, "1")
		data, err := prodsService.UpdateData(&prod, "1")
		res := data.Result
		assert.Equal(t, "Update product berhasil", res, "Result must be 'Update product berhasil'")
		assert.Nil(t, err)
	})

	t.Run("Delete Product", func(t *testing.T) {
		productsRepo.Mocks.On("Remove", "1")
		data, err := prodsService.Remove("1")
		res := data.Result
		assert.Equal(t, "Product berhasil dihapus", res, "Result must be 'Product berhasil dihapus'")
		assert.Nil(t, err)
	})
}
