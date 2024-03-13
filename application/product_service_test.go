package application

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}
