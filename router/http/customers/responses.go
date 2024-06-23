package customers

// AuthorResponse struct defines response fields
type CustomerResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ListResponse struct defines authors list response structure
type ListResponse struct {
	Data []CustomerResponse `json:"data"`
}
