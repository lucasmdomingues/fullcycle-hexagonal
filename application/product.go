package application

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gookit/validate"
)

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type Product struct {
	ID     string  `validate:"uuid4"`
	Name   string  `validate:"required"`
	Price  float64 `validate:"float"`
	Status string  `validate:"required"`
}

func NewProduct() *Product {
	return &Product{
		ID:     uuid.New().String(),
		Status: DISABLED,
	}
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("status must be enaled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("price must be greater or equal than zero")
	}

	err := validate.Struct(p).ValidateErr()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("price must be greater than zero")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("price must be zero")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetStatus() string {
	return p.Status
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}
