package repos

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type ProductRepMock struct {
	Mocks mock.Mock
}

var data = models.Product{
	Name:      "testing",
	Price:     2000,
	Id:        1,
	Category:  "1",
	CreatedAt: time.Now(),
	UpdateAt:  time.Now(),
}

var datas = models.Products{
	data,
}

func (pr *ProductRepMock) FindAll() (*models.Products, error) {
	args := pr.Mocks.Called()

	if args == nil {
		return nil, nil
	}

	return &datas, nil
}

func (pr *ProductRepMock) Save(prod *models.Product) error {
	args := pr.Mocks.Called(prod)

	if args == nil {
		return nil
	}

	return nil
}

func (pr *ProductRepMock) Remove(id string) error {
	args := pr.Mocks.Called(id)

	if args == nil {
		return nil
	}

	return nil
}

func (pr *ProductRepMock) Edit(prod *models.Product, id string) error {
	args := pr.Mocks.Called(prod, id)

	if args == nil {
		return nil
	}

	return nil
}

func (pr *ProductRepMock) FindbyPriceASC() (*models.Products, error) {
	args := pr.Mocks.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

func (pr *ProductRepMock) FindbyPriceDESC() (*models.Products, error) {
	args := pr.Mocks.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

func (pr *ProductRepMock) FindbyDateASC() (*models.Products, error) {
	args := pr.Mocks.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

func (pr *ProductRepMock) FindbyDateDESC() (*models.Products, error) {
	args := pr.Mocks.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

func (pr *ProductRepMock) FindbyCategory() (*models.Products, error) {
	args := pr.Mocks.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

func (pr *ProductRepMock) SearchProductCategory(name string) (*models.Products, error) {
	args := pr.Mocks.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}

func (pr *ProductRepMock) SearchProductName(name string) (*models.Products, error) {
	args := pr.Mocks.Called()

	if args.Get(0) == nil {
		return nil, nil
	}

	return nil, nil
}
