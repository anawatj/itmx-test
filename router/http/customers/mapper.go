package customers

import (
	domain "itmx-test/domain/customers"
)

func toResponseModel(entity *domain.Customer) *CustomerResponse {
	return &CustomerResponse{
		Id:   entity.Id,
		Name: entity.Name,
		Age:  entity.Age,
	}

}
