package product

import (
	"kholabazar/domain"
	productHandler "kholabazar/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Create(product domain.Product) (*domain.Product, error)
	Get(ID int) (*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(ID int) error
}
