package customers

import (
	domain "itmx-test/domain/customers"
)

func toDBModel(entity *domain.Customer) *Customer {
	return &Customer{
		Id:   entity.Id,
		Name: entity.Name,
		Age:  entity.Age,
	}
}

func toDomainModel(entity *Customer) *domain.Customer {
	return &domain.Customer{
		Id:   entity.Id,
		Name: entity.Name,
		Age:  entity.Age,
	}
}
