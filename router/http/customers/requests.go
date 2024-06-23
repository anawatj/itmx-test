package customers

import (
	customers "itmx-test/domain/customers"

	"github.com/gin-gonic/gin"
)

type CustomerRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Bind validates the JSON payload and returns a Author instance corresponding to the payload.
func Bind(c *gin.Context) (*customers.Customer, error) {
	var json CustomerRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	customer := &customers.Customer{
		Name: json.Name,
		Age:  json.Age,
	}

	return customer, nil
}
