package customers

// struct defines the database model for a Author.
type Customer struct {
	Id   int `gorm:"primary_key";"AUTO_INCREMENT";`
	Name string
	Age  int
}
