package application_test

import (
	"github.com/codeedu/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	uuidTest := uuid.NewV4().String()
	product.ID = uuidTest

	require.Equal(t, uuidTest, product.GetID())
	require.Equal(t, "Hello", product.GetName())
	require.Equal(t, application.DISABLED, product.GetStatus())
	require.Equal(t, 10.0, product.GetPrice())
}

func TestNewProduct(t *testing.T) {
	product := application.NewProduct()
	require.NotEmpty(t, product.GetID())
	require.Equal(t, application.DISABLED, product.GetStatus())
}
