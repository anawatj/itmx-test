package mocks

import (
	"itmx-test/domain/customers"

	"github.com/stretchr/testify/mock"
)

type CustomerRepoMock struct {
	mock.Mock
	CustomerRepo customers.CustomerRepository
}

func (m *CustomerRepoMock) GetAllCustomer() ([]customers.Customer, error) {
	arguments := m.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(error)
	} else {
		return arguments.Get(0).([]customers.Customer), nil
	}
}
func (m *CustomerRepoMock) CreateCustomer(customer *customers.Customer) (*customers.Customer, error) {
	arguments := m.Called(customer)
	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(error)
	} else {
		return arguments.Get(0).(*customers.Customer), nil
	}

}
func (m *CustomerRepoMock) UpdateCustomer(customer *customers.Customer, id int) (*customers.Customer, error) {
	arguments := m.Called(customer, id)
	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(error)
	} else {
		return arguments.Get(0).(*customers.Customer), nil
	}
}

func (m *CustomerRepoMock) GetByIdCustomer(id int) (*customers.Customer, error) {
	arguments := m.Called(id)
	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(error)
	} else {
		return arguments.Get(0).(*customers.Customer), nil
	}
}
func (m *CustomerRepoMock) DeleteCustomer(id int) error {

	arguments := m.Called(id)
	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	}
	return nil

}
func (m *CustomerRepoMock) CheckDuplicateCustomerName(name string, id int) bool {

	arguments := m.Called(name, id)
	return arguments.Get(0).(bool)
}
