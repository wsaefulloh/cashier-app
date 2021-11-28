package interfaces

import (
	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type ProductServices interface {
	FindAll() (*helpers.Response, error)
	AddData(prod *models.Product) (*helpers.Response, error)
	Remove(id string) (*helpers.Response, error)
	UpdateData(prod *models.Product, id string) (*helpers.Response, error)
	SearchbyName(name string) (*helpers.Response, error)
	SearchbyCategory(category string) (*helpers.Response, error)
	SortbyCategory() (*helpers.Response, error)
	SortbyDateASC() (*helpers.Response, error)
	SortbyDateDESC() (*helpers.Response, error)
	SortbyPriceASC() (*helpers.Response, error)
	SortbyPriceDESC() (*helpers.Response, error)
}
