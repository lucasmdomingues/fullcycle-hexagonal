package cli

import (
	"fmt"

	"github.com/lucasmdomingues/hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productID, productName string,
	productPrice float64,
	productStatus string,
) (string, error) {
	var message string

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return message, err
		}

		message = fmt.Sprintf("Product ID %s with name %s has been created with price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return message, err
		}

		_, err = service.Enable(product)
		if err != nil {
			return message, err
		}

		message = fmt.Sprintf("Product %s has been enabled", product.GetName())
	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return message, err
		}

		_, err = service.Disable(product)
		if err != nil {
			return message, err
		}

		message = fmt.Sprintf("Product %s has been disabled", product.GetName())
	default:
		product, err := service.Get(productID)
		if err != nil {
			return message, err
		}

		message = fmt.Sprintf("%+v", product)
	}

	return message, nil
}
