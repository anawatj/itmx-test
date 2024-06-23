package customers_test

import (
	"itmx-test/domain/customers"
	"itmx-test/mocks"
	"net/http"
	"testing"

	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	NotFound = "record not found"
)

func TestCreateCustomerService(t *testing.T) {
	t.Run("Test Create Customer Service success", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   0,
			Name: "Test22",
			Age:  10,
		}
		customerRepoMock.On("CreateCustomer", mock.AnythingOfType("*customers.Customer")).Return(&customer, nil)
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(true)
		service := customers.NewService(customerRepoMock)
		_, code, err := service.CreateCustomer(&customer)
		assert.Equal(t, http.StatusCreated, code)
		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
	t.Run("Test Create Customer Service required field", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   0,
			Name: "",
			Age:  -1,
		}
		customerRepoMock.On("CreateCustomer", mock.AnythingOfType("*customers.Customer")).Return(&customer, nil)
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(true)
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.CreateCustomer(&customer)
		assert.Equal(t, http.StatusBadRequest, code)

	})

	t.Run("Test Create Customer Service with duplicate customer name", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   0,
			Name: "Test22",
			Age:  10,
		}
		customerRepoMock.On("CreateCustomer", mock.AnythingOfType("*customers.Customer")).Return(&customer, nil)
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(false)
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.CreateCustomer(&customer)
		assert.Equal(t, http.StatusBadRequest, code)

	})

	t.Run("Test Create Customer Service database error", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   0,
			Name: "Test22",
			Age:  10,
		}
		customerRepoMock.On("CreateCustomer", mock.AnythingOfType("*customers.Customer")).Return(nil, errors.New("Database Error"))
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(true)
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.CreateCustomer(&customer)
		assert.Equal(t, http.StatusInternalServerError, code)

	})
}

func TestGetAllCustomerService(t *testing.T) {
	t.Run("Test Get All Customer Service success", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		var results []customers.Customer
		results = append(results, customers.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		})
		customerRepoMock.On("GetAllCustomer").Return(results, nil)

		service := customers.NewService(customerRepoMock)
		_, code, _ := service.GetAllCustomer()
		assert.Equal(t, http.StatusOK, code)

	})

	t.Run("Test Get All Customer Service not found", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		var results []customers.Customer

		customerRepoMock.On("GetAllCustomer").Return(results, nil)

		service := customers.NewService(customerRepoMock)
		_, code, _ := service.GetAllCustomer()
		assert.Equal(t, http.StatusNotFound, code)

	})

	t.Run("Test Get All Customer Service error", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customerRepoMock.On("GetAllCustomer").Return(nil, errors.New("Error"))
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.GetAllCustomer()
		assert.Equal(t, http.StatusInternalServerError, code)
	})
}

func TestGetByIdCustomerService(t *testing.T) {
	t.Run("Test Get by id Customer Service success", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)

		customer := customers.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		}
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(&customer, nil)

		service := customers.NewService(customerRepoMock)
		_, code, _ := service.GetByIdCustomer(1)
		assert.Equal(t, http.StatusOK, code)

	})
	t.Run("Test Get by id Customer Service not found", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)

		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(nil, errors.New(NotFound))

		service := customers.NewService(customerRepoMock)
		_, code, _ := service.GetByIdCustomer(1)
		assert.Equal(t, http.StatusNotFound, code)

	})

	t.Run("Test Get by id Customer Service error", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)

		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(nil, errors.New("Error"))

		service := customers.NewService(customerRepoMock)
		_, code, _ := service.GetByIdCustomer(1)
		assert.Equal(t, http.StatusInternalServerError, code)
	})
}

func TestUpdateCustomerService(t *testing.T) {
	t.Run("Test Update Customer Service success", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   1,
			Name: "Test22",
			Age:  10,
		}
		customerRepoMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(&customer, nil)
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(true)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(&customer, nil)
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.UpdateCustomer(&customer, 1)
		assert.Equal(t, http.StatusOK, code)

	})
	t.Run("Test Update Customer Service required field", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   1,
			Name: "",
			Age:  -1,
		}
		customerRepoMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(&customer, nil)
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(true)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(&customer, nil)
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.UpdateCustomer(&customer, 1)
		assert.Equal(t, http.StatusBadRequest, code)

	})

	t.Run("Test Update Customer Service with duplicate customer name", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   1,
			Name: "Test22",
			Age:  10,
		}
		customerRepoMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(&customer, nil)
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(false)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(&customer, nil)
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.UpdateCustomer(&customer, 1)
		assert.Equal(t, http.StatusBadRequest, code)

	})

	t.Run("Test Update Customer Service database error", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   1,
			Name: "Test22",
			Age:  10,
		}
		customerRepoMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(nil, errors.New("Error"))
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(true)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(&customer, nil)
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.UpdateCustomer(&customer, 1)
		assert.Equal(t, http.StatusInternalServerError, code)

	})

	t.Run("Test Update Customer Service customer not found", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   1,
			Name: "Test22",
			Age:  10,
		}
		customerRepoMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(&customer, nil)
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(true)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(nil, errors.New(NotFound))
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.UpdateCustomer(&customer, 1)
		assert.Equal(t, http.StatusNotFound, code)

	})

	t.Run("Test Update Customer Service customer get by id error", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)
		customer := customers.Customer{
			Id:   1,
			Name: "Test22",
			Age:  10,
		}
		customerRepoMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(&customer, nil)
		customerRepoMock.On("CheckDuplicateCustomerName", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(true)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(nil, errors.New("Error"))
		service := customers.NewService(customerRepoMock)
		_, code, _ := service.UpdateCustomer(&customer, 1)
		assert.Equal(t, http.StatusInternalServerError, code)

	})
}

func TestDeleteCustomerService(t *testing.T) {
	t.Run("Test Delete Customer Service success", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)

		customer := customers.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		}
		customerRepoMock.On("DeleteCustomer", mock.AnythingOfType("int")).Return(nil)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(&customer, nil)

		service := customers.NewService(customerRepoMock)
		code, _ := service.DeleteCustomer(1)
		assert.Equal(t, http.StatusOK, code)

	})
	t.Run("Test Delete Customer Service not found", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)

		customerRepoMock.On("DeleteCustomer", mock.AnythingOfType("int")).Return(nil)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(nil, errors.New(NotFound))

		service := customers.NewService(customerRepoMock)
		code, _ := service.DeleteCustomer(1)
		assert.Equal(t, http.StatusNotFound, code)

	})

	t.Run("Test Delete Customer Service get by id error", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)

		customerRepoMock.On("DeleteCustomer", mock.AnythingOfType("int")).Return(nil)
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(nil, errors.New("Error"))

		service := customers.NewService(customerRepoMock)
		code, _ := service.DeleteCustomer(1)
		assert.Equal(t, http.StatusInternalServerError, code)
	})

	t.Run("Test Delete Customer Service error", func(t *testing.T) {
		customerRepoMock := new(mocks.CustomerRepoMock)

		customer := customers.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		}
		customerRepoMock.On("DeleteCustomer", mock.AnythingOfType("int")).Return(errors.New("Error"))
		customerRepoMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(&customer, nil)

		service := customers.NewService(customerRepoMock)
		code, _ := service.DeleteCustomer(1)
		assert.Equal(t, http.StatusInternalServerError, code)

	})
}
