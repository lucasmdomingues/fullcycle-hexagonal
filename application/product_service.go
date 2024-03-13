package application

import "log"

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		log.Print("failed to get product", err)
		return nil, err
	}

	return product, nil
}
