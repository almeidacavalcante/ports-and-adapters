package cli_test

import (
	"fmt"
	"github.com/almeidacavalcante/ports-and-adapters/adapters/cli"
	mock_application "github.com/almeidacavalcante/ports-and-adapters/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 10.0
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus)
	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(serviceMock, "enable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(serviceMock, "disable", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nStatus: %s\nPrice: %f",
		productId, productName, productStatus, productPrice)
	result, err = cli.Run(serviceMock, "", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
