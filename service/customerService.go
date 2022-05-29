package service

import "banking/domain"

// CustomerService service port is interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// DefaultCustomerService business logic
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()

}

//instantiate the default customer service

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
