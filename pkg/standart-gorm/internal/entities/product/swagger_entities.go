package product

type RequestBodyToCreateOrUpdateProduct struct {
	Name string `json:"name"`
}

type ResponseCreateProduct struct {
	ID string `json:"id"`
}

type ResponseListAllProduct struct {
	Status int       `json:"status"`
	Result []Product `json:"result"`
}

type ResponseListOneProduct struct {
	Status int     `json:"status"`
	Result Product `json:"result"`
}
