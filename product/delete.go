package product


func (svc *service) Delete(ID int) error {
	return svc.productRepo.Delete(ID)
}
