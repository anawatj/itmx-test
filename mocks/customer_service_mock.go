package mocks

import (
	"itmx-test/domain/customers"

	"github.com/stretchr/testify/mock"
)

type CustomerServiceMock struct {
	mock.Mock
	CustomerService customers.CustomerService
}

func (m *CustomerServiceMock) CreateCustomer(customer *customers.Customer) (*customers.Customer, int, error) {
	arguments := m.Called(customer)
	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(int), arguments.Get(2).(error)
	}
	return arguments.Get(0).(*customers.Customer), arguments.Get(1).(int), nil
}

func (m *CustomerServiceMock) UpdateCustomer(customer *customers.Customer, id int) (*customers.Customer, int, error) {
	arguments := m.Called(customer, id)
	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(int), arguments.Get(2).(error)
	}
	return arguments.Get(0).(*customers.Customer), arguments.Get(1).(int), nil
}
func (m *CustomerServiceMock) GetAllCustomer() ([]customers.Customer, int, error) {
	arguments := m.Called()
	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(int), arguments.Get(2).(error)
	}
	return arguments.Get(0).([]customers.Customer), arguments.Get(1).(int), nil
}

func (m *CustomerServiceMock) GetByIdCustomer(id int) (*customers.Customer, int, error) {
	arguments := m.Called(id)
	if arguments.Get(0) == nil {
		return nil, arguments.Get(1).(int), arguments.Get(2).(error)
	}
	return arguments.Get(0).(*customers.Customer), arguments.Get(1).(int), nil
}

func (m *CustomerServiceMock) DeleteCustomer(id int) (int, error) {
	arguments := m.Called(id)
	if arguments.Get(1) != nil {
		return arguments.Get(0).(int), arguments.Get(1).(error)
	}
	return arguments.Get(0).(int), nil
}
