package customers

import (
	domain "itmx-test/domain/customers"

	"github.com/jinzhu/gorm"
)

const (
	NotFound = "record not found"
)

type Store struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Store {
	db.AutoMigrate(&Customer{})

	return &Store{
		DB: db,
	}
}

func (s *Store) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	entity := toDBModel(customer)
	err := s.DB.Create(entity).Error
	return toDomainModel(entity), err
}
func (s *Store) GetAllCustomer() ([]domain.Customer, error) {
	var results []Customer
	err := s.DB.Find(&results).Error
	var ret []domain.Customer
	if err == nil {
		var customers = make([]domain.Customer, len(results))
		for i, element := range results {
			customers[i] = *toDomainModel(&element)
		}
		ret = customers
	}
	return ret, err
}
func (s *Store) GetByIdCustomer(id int) (*domain.Customer, error) {
	result := &Customer{}
	err := s.DB.Where("id = ?", id).First(result).Error
	return toDomainModel(result), err
}
func (s *Store) UpdateCustomer(customer *domain.Customer, id int) (*domain.Customer, error) {
	entity := toDBModel(customer)
	err := s.DB.Save(entity).Error

	return toDomainModel(entity), err
}
func (s *Store) DeleteCustomer(id int) error {
	err := s.DB.Delete(&Customer{Id: id}).Error
	return err
}
func (s *Store) CheckDuplicateCustomerName(name string, id int) bool {
	result := &Customer{}
	err := s.DB.Where("name=? AND id<>?", name, id).First(result).Error
	if err != nil {
		return err.Error() == NotFound
	} else {
		return false
	}
}
