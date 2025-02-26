package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ID: "1", Name: "John", City: "a", ZipCode: "1", DateOfBirth: "2000-01-01", Status: "1"},
		{ID: "2", Name: "Ron", City: "b", ZipCode: "2", DateOfBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
