package app

import "errors"

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED       = "disabled"
	ENABLED        = "enabled"
	ENABLED_ERROR  = "the price must be greater than zero to enable the product"
	DISABLED_ERROR = "the price must be zero in order to have the product disabled"
	STATUS_ERROR   = "the product status must be ENABLED or DISABLED"
	PRICE_ERROR    = "the price must be greater than zero to enable the product"
)

type Product struct {
	ID     string
	Name   string
	Status string
	Price  float64
}

func (product *Product) IsValid() (bool, error) {
	if product.Status != ENABLED && product.Status != DISABLED {
		return false, errors.New(STATUS_ERROR)
	}

	if product.Price <= 0 {
		return false, errors.New(PRICE_ERROR)
	}

	return true, nil
}

func (product *Product) Enable() error {
	if product.Price > 0 {
		product.Status = ENABLED

		return nil
	}

	return errors.New(ENABLED_ERROR)
}

func (product *Product) Disable() error {
	if product.Price == 0 {
		product.Status = DISABLED

		return nil
	}

	return errors.New(DISABLED_ERROR)
}

func (product *Product) GetID() string {
	return product.ID
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) GetStatus() string {
	return product.Status
}

func (product *Product) GetPrice() float64 {
	return product.Price
}
