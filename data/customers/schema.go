package customers

type Customer struct {
	Id   int `gorm:"primary_key";"AUTO_INCREMENT";`
	Name string
	Age  int
}
