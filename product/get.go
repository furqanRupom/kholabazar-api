package product

import "kholabazar/domain"

func (svc *service) Get(ID int) (*domain.Product, error) {
	return svc.productRepo.Get(ID)

}
