package product

import (
	"kholabazar/domain"
)

type Service interface {
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Create(product domain.Product) (*domain.Product, error)
	Get(ID int) (*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(ID int) error
}
