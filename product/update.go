package product

import "kholabazar/domain"

func (svc *service) Update(product domain.Product) (*domain.Product, error) {
	return svc.productRepo.Update(product)
	
}
