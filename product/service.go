package product

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}
