package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {

	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Charlespogz", "Davao City", "8000", "Nov 20 1970", "1"},
		{"1002", "Sabsab", "Davao City", "9000", "may 21 2016", "1"},
		{"1002", "chaz", "Davao City", "7000", "Nov 20 1970", "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
