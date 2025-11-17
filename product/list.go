package product

import "kholabazar/domain"

func (svc *service) List(page,limit int64) ([]*domain.Product, error) {
	return svc.productRepo.List(page,limit)
}
