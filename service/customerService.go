package service

import "github.com/manjusharami/bankApp/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepositry

}
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
    return s.repo.FindAll()
}
 

func NewCustomerService(repository domain.CustomerRepositry) DefaultCustomerService{
	return DefaultCustomerService{repo: repository}
}