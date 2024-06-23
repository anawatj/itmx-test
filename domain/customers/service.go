package customers

import (
	"errors"
	"net/http"
	"strings"
)

const (
	NotFound = "record not found"
)

type CustomerService interface {
	CreateCustomer(*Customer) (*Customer, int, error)
	GetAllCustomer() ([]Customer, int, error)
	GetByIdCustomer(int) (*Customer, int, error)
	UpdateCustomer(*Customer, int) (*Customer, int, error)
	DeleteCustomer(int) (int, error)
}
type Service struct {
	Repository CustomerRepository
}

func (svc *Service) CreateCustomer(customer *Customer) (*Customer, int, error) {
	var dataErrors []string
	if len(customer.Name) == 0 {
		dataErrors = append(dataErrors, "Customer name is required")
	}
	if customer.Age < 0 {
		dataErrors = append(dataErrors, "Age is required")
	}
	if len(dataErrors) > 0 {
		return nil, http.StatusBadRequest, errors.New(strings.Join(dataErrors, ","))
	}
	isNotFound := svc.Repository.CheckDuplicateCustomerName(customer.Name, 0)
	if !isNotFound {
		return nil, http.StatusBadRequest, errors.New("Customer Name is Dduplicate")
	}

	ret, err := svc.Repository.CreateCustomer(customer)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(err.Error())
	}
	return ret, http.StatusCreated, nil
}
func (svc *Service) GetAllCustomer() ([]Customer, int, error) {
	ret, err := svc.Repository.GetAllCustomer()
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(err.Error())
	}
	if len(ret) == 0 {
		return nil, http.StatusNotFound, errors.New("Customer Not Found")
	}
	return ret, http.StatusOK, nil
}
func (svc *Service) GetByIdCustomer(id int) (*Customer, int, error) {
	ret, err := svc.Repository.GetByIdCustomer(id)
	if err != nil {
		if err.Error() == NotFound {
			return nil, http.StatusNotFound, errors.New("Customer Not Found")
		} else {
			return nil, http.StatusInternalServerError, errors.New(err.Error())
		}
	}

	return ret, http.StatusOK, nil
}
func (svc *Service) UpdateCustomer(customer *Customer, id int) (*Customer, int, error) {
	var dataErrors []string
	if len(customer.Name) == 0 {
		dataErrors = append(dataErrors, "Customer name is required")
	}
	if customer.Age < 0 {
		dataErrors = append(dataErrors, "Age is required")
	}
	if len(dataErrors) > 0 {
		return nil, http.StatusBadRequest, errors.New(strings.Join(dataErrors, ","))
	}
	isNotFound := svc.Repository.CheckDuplicateCustomerName(customer.Name, id)
	if !isNotFound {
		return nil, http.StatusBadRequest, errors.New("Customer Name is Dduplicate")
	}
	customerDb, err := svc.Repository.GetByIdCustomer(id)
	if err != nil {
		if err.Error() == NotFound {
			return nil, http.StatusNotFound, errors.New("Customer Not Found")
		} else {
			return nil, http.StatusInternalServerError, errors.New(err.Error())
		}
	}

	customerDb.Name = customer.Name
	customerDb.Age = customer.Age
	ret, err := svc.Repository.UpdateCustomer(customerDb, id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return ret, http.StatusOK, nil
}

func (svc *Service) DeleteCustomer(id int) (int, error) {

	_, err := svc.Repository.GetByIdCustomer(id)
	if err != nil {
		if err.Error() == NotFound {
			return http.StatusNotFound, errors.New("Customer Not Found")
		} else {
			return http.StatusInternalServerError, errors.New(err.Error())
		}

	}

	err = svc.Repository.DeleteCustomer(id)
	if err != nil {
		return http.StatusInternalServerError, errors.New(err.Error())
	}

	return http.StatusOK, nil
}
func NewService(repository CustomerRepository) *Service {
	return &Service{Repository: repository}
}
