package customers_test

import (
	"database/sql"
	"errors"
	"itmx-test/data/customers"
	domains "itmx-test/domain/customers"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type v1Suite struct {
	db       *gorm.DB
	mock     sqlmock.Sqlmock
	customer domains.Customer
}

func TestCreateCustomerSuccess(t *testing.T) {
	s := &v1Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("sqlite3", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	customer := domains.Customer{
		Id:   0,
		Name: "Test",
		Age:  10,
	}
	store := &customers.Store{
		DB: s.db,
	}
	customers.New(s.db)
	s.mock.ExpectBegin()
	store.CreateCustomer(&customer)
	s.mock.ExpectCommit()
	assert.Equal(t, "Test", customer.Name)
	assert.Equal(t, 10, customer.Age)

	defer db.Close()
}

func TestGetCustomerByIdSuccess(t *testing.T) {
	s := &v1Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("sqlite3", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	customer := domains.Customer{
		Id:   0,
		Name: "Test",
		Age:  10,
	}
	store := &customers.Store{
		DB: s.db,
	}

	s.mock.ExpectBegin()
	store.CreateCustomer(&customer)
	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(customer.Id, customer.Name, customer.Age)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.GetByIdCustomer(1)
	s.mock.ExpectCommit()
	defer db.Close()
}

func TestGetAllCustomerSuccess(t *testing.T) {
	s := &v1Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("sqlite3", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	customer := domains.Customer{
		Id:   0,
		Name: "Test",
		Age:  10,
	}
	store := &customers.Store{
		DB: s.db,
	}
	s.mock.ExpectBegin()
	store.CreateCustomer(&customer)
	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(customer.Id, customer.Name, customer.Age)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.GetAllCustomer()
	s.mock.ExpectCommit()
	defer db.Close()
}
func TestUpdateCustomerSuccess(t *testing.T) {
	s := &v1Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("sqlite3", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	customer := domains.Customer{
		Id:   0,
		Name: "Test",
		Age:  10,
	}
	store := &customers.Store{
		DB: s.db,
	}

	s.mock.ExpectBegin()
	store.CreateCustomer(&customer)
	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(customer.Id, customer.Name, customer.Age)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.GetByIdCustomer(1)
	customer.Name = "Test22"
	customer.Age = 45
	store.UpdateCustomer(&customer, customer.Id)
	s.mock.ExpectCommit()
	defer db.Close()
}
func TestDeleteCustomerSuccess(t *testing.T) {
	s := &v1Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("sqlite3", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	customer := domains.Customer{
		Id:   0,
		Name: "Test",
		Age:  10,
	}
	store := &customers.Store{
		DB: s.db,
	}

	s.mock.ExpectBegin()
	store.CreateCustomer(&customer)
	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(customer.Id, customer.Name, customer.Age)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.GetByIdCustomer(1)
	store.DeleteCustomer(1)
	s.mock.ExpectCommit()
	defer db.Close()
}

func TestCheckDuplicateCustomerNameFound(t *testing.T) {
	t.Run("it success", func(t *testing.T) {
		s := &v1Suite{}
		var (
			db  *sql.DB
			err error
		)

		db, s.mock, err = sqlmock.New()

		if err != nil {
			t.Errorf("Failed to open mock sql db, got error: %v", err)
		}

		if db == nil {
			t.Error("mock db is null")
		}

		if s.mock == nil {
			t.Error("sqlmock is null")
		}

		s.db, err = gorm.Open("sqlite3", db)
		if err != nil {
			t.Errorf("Failed to open gorm db, got error: %v", err)
		}

		if s.db == nil {
			t.Error("gorm db is null")
		}
		customer := domains.Customer{
			Id:   0,
			Name: "Test",
			Age:  10,
		}
		store := &customers.Store{
			DB: s.db,
		}

		s.mock.ExpectBegin()
		store.CreateCustomer(&customer)
		rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(customer.Id, customer.Name, customer.Age)
		s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
		store.CheckDuplicateCustomerName("Test", 0)
		s.mock.ExpectCommit()
		defer db.Close()
	})
	t.Run("it errors", func(t *testing.T) {
		s := &v1Suite{}
		var (
			db  *sql.DB
			err error
		)

		db, s.mock, err = sqlmock.New()

		if err != nil {
			t.Errorf("Failed to open mock sql db, got error: %v", err)
		}

		if db == nil {
			t.Error("mock db is null")
		}

		if s.mock == nil {
			t.Error("sqlmock is null")
		}

		s.db, err = gorm.Open("sqlite3", db)
		if err != nil {
			t.Errorf("Failed to open gorm db, got error: %v", err)
		}

		if s.db == nil {
			t.Error("gorm db is null")
		}
		customer := domains.Customer{
			Id:   0,
			Name: "Test",
			Age:  10,
		}
		store := &customers.Store{
			DB: s.db,
		}

		s.mock.ExpectBegin()
		store.CreateCustomer(&customer)
		rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(customer.Id, customer.Name, customer.Age)
		s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
		store.DB.AddError(errors.New("Database Error"))
		store.CheckDuplicateCustomerName("Test", 0)

		s.mock.ExpectCommit()
		defer db.Close()
	})

}

func TestCheckDuplicateCustomerNameNotFound(t *testing.T) {
	s := &v1Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	s.db, err = gorm.Open("sqlite3", db)
	if err != nil {
		t.Errorf("Failed to open gorm db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}
	customer := domains.Customer{
		Id:   0,
		Name: "Test",
		Age:  10,
	}
	store := &customers.Store{
		DB: s.db,
	}

	s.mock.ExpectBegin()
	store.CreateCustomer(&customer)
	rows := sqlmock.NewRows([]string{"Id", "Name", "Age"}).AddRow(customer.Id, customer.Name, customer.Age)
	s.mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	store.CheckDuplicateCustomerName("Test11", 1)
	s.mock.ExpectCommit()
	defer db.Close()
}
