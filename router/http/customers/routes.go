package customers

import (
	"fmt"
	customers "itmx-test/domain/customers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	Service customers.CustomerService
}

func (handler *CustomerHandler) GetAllCustomer(c *gin.Context) {
	results, code, err := handler.Service.GetAllCustomer()
	if results == nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}

	var responseItems = make([]CustomerResponse, len(results))

	for i, element := range results {
		responseItems[i] = *toResponseModel(&element)
	}

	response := &ListResponse{
		Data: responseItems,
	}

	c.JSON(code, response)
}
func (handler *CustomerHandler) CreateCustomer(c *gin.Context) {
	customer, err := Bind(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newCustomer, code, err := handler.Service.CreateCustomer(customer)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}

	c.JSON(code, *toResponseModel(newCustomer))
}

func (handler *CustomerHandler) GetByIdCustomer(c *gin.Context) {
	id := c.Param("customerId")
	i, err := strconv.Atoi(id)
	result, code, err := handler.Service.GetByIdCustomer(i)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}

	c.JSON(code, *toResponseModel(result))
}

func (handler *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id := c.Param("customerId")
	i, err := strconv.Atoi(id)
	customer, err := Bind(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedCustomer, code, err := handler.Service.UpdateCustomer(customer, i)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}
	fmt.Println(updatedCustomer)
	c.JSON(code, toResponseModel(updatedCustomer))
}
func (handler *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("customerId")
	i, err := strconv.Atoi(id)
	code, err := handler.Service.DeleteCustomer(i)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}

	c.JSON(code, "Success")
}

func NewRoutesFactory(group *gin.RouterGroup) func(service customers.CustomerService) {
	customerRoutesFactory := func(service customers.CustomerService) {
		handler := CustomerHandler{Service: service}
		group.GET("/", handler.GetAllCustomer)
		group.POST("/", handler.CreateCustomer)
		group.GET("/:customerId", handler.GetByIdCustomer)
		group.PUT("/:customerId", handler.UpdateCustomer)
		group.DELETE("/:customerId", handler.DeleteCustomer)
	}

	return customerRoutesFactory
}
