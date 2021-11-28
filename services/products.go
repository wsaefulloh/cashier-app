package services

import (
	"fmt"

	"github.com/wsaefulloh/go-solid-principle/helpers"
	"github.com/wsaefulloh/go-solid-principle/interfaces"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type product_service struct {
	rep interfaces.RepoProduct
}

func ProductServices(reps interfaces.RepoProduct) *product_service {
	return &product_service{reps}
}

func (pr *product_service) FindAll() (*helpers.Response, error) {
	data, err := pr.rep.FindAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}

func (pr *product_service) AddData(prod *models.Product) (*helpers.Response, error) {
	error := pr.rep.Save(prod)

	if error != nil {
		return nil, error
	}

	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  "Product berhasil ditambahkan",
	}

	return &respon, nil
}

func (pr *product_service) UpdateData(prod *models.Product, id string) (*helpers.Response, error) {
	error := pr.rep.Edit(prod, id)

	if error != nil {
		return nil, error
	}

	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  "Update product berhasil",
	}

	return &respon, nil
}

func (pr *product_service) Remove(id string) (*helpers.Response, error) {
	error := pr.rep.Remove(id)

	if error != nil {
		return nil, error
	}

	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  "Product berhasil dihapus",
	}

	return &respon, nil
}

func (pr *product_service) SearchbyName(name string) (*helpers.Response, error) {
	data, err := pr.rep.SearchProductName(name)
	if err != nil {
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}

func (pr *product_service) SearchbyCategory(category string) (*helpers.Response, error) {
	data, err := pr.rep.SearchProductCategory(category)
	if err != nil {
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}

func (pr *product_service) SortbyCategory() (*helpers.Response, error) {
	data, err := pr.rep.FindbyCategory()
	if err != nil {
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}

func (pr *product_service) SortbyDateASC() (*helpers.Response, error) {
	data, err := pr.rep.FindbyDateASC()
	if err != nil {
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}

func (pr *product_service) SortbyDateDESC() (*helpers.Response, error) {
	data, err := pr.rep.FindbyDateDESC()
	if err != nil {
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}

func (pr *product_service) SortbyPriceASC() (*helpers.Response, error) {
	data, err := pr.rep.FindbyPriceASC()
	if err != nil {
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}

func (pr *product_service) SortbyPriceDESC() (*helpers.Response, error) {
	data, err := pr.rep.FindbyPriceDESC()
	if err != nil {
		return nil, err
	}
	respon := helpers.Response{
		Status:  200,
		IsError: false,
		Result:  data,
	}

	return &respon, nil
}
