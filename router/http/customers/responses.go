package customers

type CustomerResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ListResponse struct {
	Data []CustomerResponse `json:"data"`
}
