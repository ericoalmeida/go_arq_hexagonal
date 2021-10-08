package app_test

import (
	"testing"

	"github.com/ericoalmeida/go_arq_hexagonal/app"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

const (
	PRODUCT_PRICE_POSITIVE = 5
	PRODUCT_PRICE_ZERO     = 0
	PRODUCT_PRICE_NEGATIVE = -5
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}

	product.Name = "Product"
	product.Status = app.DISABLED
	product.Price = PRODUCT_PRICE_POSITIVE

	err := product.Enable()
	require.Nil(t, err)

	product.Price = PRODUCT_PRICE_ZERO

	err = product.Enable()
	require.Equal(t, app.ENABLED_ERROR, err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}

	product.Name = "Product"
	product.Status = app.ENABLED
	product.Price = PRODUCT_PRICE_ZERO

	err := product.Disable()
	require.Nil(t, err)

	product.Price = PRODUCT_PRICE_POSITIVE

	err = product.Disable()
	require.Equal(t, app.DISABLED_ERROR, err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Product"
	product.Status = app.DISABLED
	product.Price = PRODUCT_PRICE_POSITIVE

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, app.STATUS_ERROR, err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = PRODUCT_PRICE_NEGATIVE
	_, err = product.IsValid()
	require.Equal(t, app.PRICE_ERROR, err.Error())
}

func TestProduct_GetID(t *testing.T) {
	var productUuid string = uuid.NewV4().String()

	product := app.Product{}
	product.ID = productUuid

	require.Equal(t, product.GetID(), productUuid)
}

func TestProduct_GetName(t *testing.T) {
	var productName string = "ProductName"

	product := app.Product{}
	product.Name = productName

	require.Equal(t, product.GetName(), productName)
}

func TestProduct_GetPrice(t *testing.T) {
	var productPrice float64 = PRODUCT_PRICE_POSITIVE

	product := app.Product{}
	product.Price = productPrice

	require.Equal(t, product.GetPrice(), productPrice)
}

func TestProduct_GetStatus(t *testing.T) {
	var productStatus string = app.ENABLED

	product := app.Product{}
	product.Status = productStatus

	require.Equal(t, product.GetStatus(), productStatus)
}
