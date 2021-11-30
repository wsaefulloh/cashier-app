package services

import (
	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type ProductServiceMock struct {
	Mocks mock.Mock
}

var data = helpers.Response{
	Status: 200,
}

func (pr *ProductServiceMock) FindAll() (*helpers.Response, error) {
	args := pr.Mocks.Called()

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) AddData(prod *models.Product) (*helpers.Response, error) {

	prod = &models.Product{
		Name:     "testing",
		Price:    2000,
		Category: "1",
	}

	args := pr.Mocks.Called(prod)

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) Remove(id string) (*helpers.Response, error) {
	args := pr.Mocks.Called(id)

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) UpdateData(prod *models.Product, id string) (*helpers.Response, error) {
	prod = &models.Product{
		Name:     "testing",
		Price:    2000,
		Category: "1",
	}

	args := pr.Mocks.Called(prod, id)

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) SearchbyName(name string) (*helpers.Response, error) {
	args := pr.Mocks.Called(name)

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) SearchbyCategory(category string) (*helpers.Response, error) {
	args := pr.Mocks.Called(category)

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) SortbyCategory() (*helpers.Response, error) {
	args := pr.Mocks.Called()

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) SortbyDateASC() (*helpers.Response, error) {
	args := pr.Mocks.Called()

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) SortbyDateDESC() (*helpers.Response, error) {
	args := pr.Mocks.Called()

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) SortbyPriceASC() (*helpers.Response, error) {
	args := pr.Mocks.Called()

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}

func (pr *ProductServiceMock) SortbyPriceDESC() (*helpers.Response, error) {
	args := pr.Mocks.Called()

	if args == nil {
		data.IsError = true
		return nil, nil
	}
	data.IsError = false
	return &data, nil
}
