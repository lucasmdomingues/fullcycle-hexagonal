package cli_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/lucasmdomingues/hexagonal/adapters/cli"
	"github.com/lucasmdomingues/hexagonal/application"
	mock_application "github.com/lucasmdomingues/hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 25.99
	productStatus := application.ENABLED
	productID := uuid.New().String()

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	// CREATE
	message, err := cli.Run(service, "create", productID, productName, productPrice, productStatus)
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("Product ID %s with name %s has been created with price %f and status %s",
		productMock.GetID(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus()), message)

	// ENABLE
	message, err = cli.Run(service, "enable", productID, productName, productPrice, productStatus)
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("Product %s has been enabled", productMock.GetName()), message)

	// DISABLE
	message, err = cli.Run(service, "disable", productID, productName, productPrice, productStatus)
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("Product %s has been disabled", productMock.GetName()), message)

	// GET
	message, err = cli.Run(service, "get", productID, productName, productPrice, productStatus)
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("Product ID: %s Name:%s Price:%f Status:%s", productMock.GetID(),
		productMock.GetName(), productMock.GetPrice(), productMock.GetStatus()), message)
}
