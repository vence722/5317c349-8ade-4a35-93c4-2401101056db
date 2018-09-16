package models

// error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// create an error response
func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Error: err.Error()}
}

// create order response
type CreateOrderResponse struct {
	ID       string  `json:"id"`
	Distance float64 `json:"distance"`
	Status   string  `json:"status"`
}

// create an create order response
func NewCreateOrderResponse(order *Order) *CreateOrderResponse {
	return &CreateOrderResponse{
		ID:       order.ID,
		Distance: order.Distance,
		Status:   order.Status,
	}
}
