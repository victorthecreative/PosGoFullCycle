package entity

import (
	"errors"
	"github.com/victorthecreative/PosGoFullCycle/API/pkg/entity"
	"time"
)

var (
	ErrIDIsRequerid    = errors.New("ID is requerid")
	ErrIDIsInvalid     = errors.New("ID is Invalid")
	ErrNameIsRequerid  = errors.New("Name is requerid")
	ErrPriceIsRequerid = errors.New("Price is requerid")
	ErrInvalidPrice    = errors.New("Invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequerid
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDIsInvalid
	}
	if p.Name == "" {
		return ErrNameIsRequerid
	}
	if p.Price == 0 {
		return ErrPriceIsRequerid
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
