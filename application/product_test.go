package application

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := Product{
		Name:   "Hello",
		Price:  10,
		Status: DISABLED,
	}

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.EqualError(t, err, "price must be greater than zero")
}

func TestProduct_Disable(t *testing.T) {
	product := Product{
		Name:   "Hello",
		Price:  0,
		Status: DISABLED,
	}

	err := product.Disable()
	require.NoError(t, err)

	product.Price = 10
	err = product.Disable()
	require.EqualError(t, err, "price must be zero")
}

func TestProduct_IsValid(t *testing.T) {
	product := Product{
		ID:     uuid.New().String(),
		Name:   "Hello",
		Price:  10,
		Status: DISABLED,
	}

	_, err := product.IsValid()
	require.NoError(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.EqualError(t, err, "status must be enaled or disabled")

	product.Status = ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.EqualError(t, err, "price must be greater or equal than zero")

	product.ID = "asd"
	product.Price = 10
	_, err = product.IsValid()
	require.EqualError(t, err, "uuid4: ID value should be a UUID4 string")
}

func TestProduct_GetID(t *testing.T) {
	productID := uuid.New().String()
	product := Product{
		ID: productID,
	}

	got := product.GetID()
	require.Equal(t, productID, got)
}

func TestProduct_GetName(t *testing.T) {
	productName := "test"
	product := Product{
		Name: productName,
	}

	got := product.GetName()
	require.Equal(t, productName, got)
}

func TestProduct_GetProduct(t *testing.T) {
	productPrice := 10.0
	product := Product{
		Price: productPrice,
	}

	got := product.GetPrice()
	require.Equal(t, productPrice, got)
}

func TestProduct_GetStatus(t *testing.T) {
	productStatus := ENABLED
	product := Product{
		Status: productStatus,
	}

	got := product.GetStatus()
	require.Equal(t, productStatus, got)
}
