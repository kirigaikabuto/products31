package products31

type CreateProductCommand struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type HttpError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}