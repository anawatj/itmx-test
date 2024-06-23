package customers_test

import (
	"encoding/json"
	"errors"
	domains "itmx-test/domain/customers"
	"itmx-test/mocks"
	"itmx-test/router/http/customers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomerApi(t *testing.T) {
	t.Run("test create customer api success", func(t *testing.T) {
		var mockCustomer = customers.CustomerRequest{
			Name: "Test",
			Age:  10,
		}
		customer := domains.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		}
		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("CreateCustomer", mock.AnythingOfType("*customers.Customer")).Return(&customer, http.StatusCreated, nil)
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.POST("/customers", handler.CreateCustomer)
		body, err := json.Marshal(mockCustomer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusCreated, rec.Code)

		})
	})

	t.Run("test create customer api error", func(t *testing.T) {
		var mockCustomer = customers.CustomerRequest{
			Name: "Test",
			Age:  10,
		}

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("CreateCustomer", mock.AnythingOfType("*customers.Customer")).Return(nil, http.StatusInternalServerError, errors.New("Database Error"))
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.POST("/customers", handler.CreateCustomer)
		body, err := json.Marshal(mockCustomer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

		})
	})

	t.Run("test create customer api json error", func(t *testing.T) {
		var mockCustomer = `{
			"Name":"Test",
			"Age":10,
			"Status":"Complete"
		
		}`

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("CreateCustomer", mock.AnythingOfType("*customers.Customer")).Return(nil, http.StatusInternalServerError, errors.New("Database Error"))
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.POST("/customers", handler.CreateCustomer)
		body, err := json.Marshal(mockCustomer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)

		})
	})
}

func TestGetAllCustomerApi(t *testing.T) {
	t.Run("test get all customer api success", func(t *testing.T) {

		ret := []domains.Customer{}

		ret = append(ret, domains.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		})

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("GetAllCustomer").Return(ret, http.StatusOK, nil)
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.GET("/customers", handler.GetAllCustomer)

		req := httptest.NewRequest(http.MethodGet, "/customers", nil)
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)

		})
	})

	t.Run("test get all customer api error", func(t *testing.T) {

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("GetAllCustomer").Return(nil, http.StatusInternalServerError, errors.New("Database Error"))
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.GET("/customers", handler.GetAllCustomer)

		req := httptest.NewRequest(http.MethodGet, "/customers", nil)
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

		})
	})
}
func TestGetByIdCustomerApi(t *testing.T) {
	t.Run("test get by id customer api success", func(t *testing.T) {

		ret := domains.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		}

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(&ret, http.StatusOK, nil)
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.GET("/customers/:customerId", handler.GetByIdCustomer)

		req := httptest.NewRequest(http.MethodGet, "/customers/1", nil)
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)

		})
	})

	t.Run("test get by id customer api error", func(t *testing.T) {

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("GetByIdCustomer", mock.AnythingOfType("int")).Return(nil, http.StatusInternalServerError, errors.New("Database Error"))
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.GET("/customers/:customerId", handler.GetByIdCustomer)

		req := httptest.NewRequest(http.MethodGet, "/customers/1", nil)
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

		})
	})
}

func TestUpdateCustomerApi(t *testing.T) {
	t.Run("test update customer api success", func(t *testing.T) {
		var mockCustomer = customers.CustomerRequest{
			Name: "Test",
			Age:  10,
		}
		customer := domains.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		}
		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(&customer, http.StatusOK, nil)
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.PUT("/customers/:customerId", handler.UpdateCustomer)
		body, err := json.Marshal(mockCustomer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/customers/1", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)

		})
	})

	t.Run("test update customer api error", func(t *testing.T) {
		var mockCustomer = customers.CustomerRequest{
			Name: "Test",
			Age:  10,
		}

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(nil, http.StatusInternalServerError, errors.New("Database Error"))
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.PUT("/customers/:customerId", handler.UpdateCustomer)
		body, err := json.Marshal(mockCustomer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/customers/1", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

		})
	})

	t.Run("test update customer api json error", func(t *testing.T) {
		var mockCustomer = `{
			"Name":"Test",
			"Age":10,
			"Status":"Complete"
		
		}`

		customer := domains.Customer{
			Id:   1,
			Name: "Test",
			Age:  10,
		}
		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("UpdateCustomer", mock.AnythingOfType("*customers.Customer"), mock.AnythingOfType("int")).Return(&customer, http.StatusOK, nil)
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.PUT("/customers/:customerId", handler.UpdateCustomer)
		body, err := json.Marshal(mockCustomer)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/customers/1", strings.NewReader(string(body)))
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)

		})
	})
}

func TestDeleteCustomerApi(t *testing.T) {
	t.Run("test delete customer api success", func(t *testing.T) {

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("DeleteCustomer", mock.AnythingOfType("int")).Return(http.StatusOK, nil)
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.DELETE("/customers/:customerId", handler.DeleteCustomer)

		req := httptest.NewRequest(http.MethodDelete, "/customers/1", nil)
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)

		})
	})

	t.Run("test delete customer api error", func(t *testing.T) {

		customerServiceMock := new(mocks.CustomerServiceMock)
		customerServiceMock.On("DeleteCustomer", mock.AnythingOfType("int")).Return(http.StatusInternalServerError, errors.New("Database Error"))
		gin := gin.New()
		rec := httptest.NewRecorder()
		handler := customers.CustomerHandler{
			Service: customerServiceMock,
		}
		gin.DELETE("/customers/:customerId", handler.DeleteCustomer)

		req := httptest.NewRequest(http.MethodDelete, "/customers/1", nil)
		gin.ServeHTTP(rec, req)

		t.Run("test status code and response body", func(t *testing.T) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

		})
	})
}
