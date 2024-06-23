package customers

type CustomerRepository interface {
	CreateCustomer(*Customer) (*Customer, error)
	GetAllCustomer() ([]Customer, error)
	GetByIdCustomer(int) (*Customer, error)
	UpdateCustomer(*Customer, int) (*Customer, error)
	DeleteCustomer(int) error
	CheckDuplicateCustomerName(string, int) bool
}
