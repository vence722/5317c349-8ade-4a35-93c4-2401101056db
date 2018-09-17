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

// take order response
type TakeOrderResponse struct {
	Status string `json:"status"`
}

// create an take order response
func NewTakeOrderResponse() *TakeOrderResponse {
	return &TakeOrderResponse{
		Status: "SUCCESS",
	}
}

type ListOrderEntry struct {
	ID       string  `json:"id"`
	Distance float64 `json:"distance"`
	Status   string  `json:"status"`
}

type ListOrdersResponse []*ListOrderEntry

func NewListOrdersResponse(orders []*Order) ListOrdersResponse {
	listOrdersResp := ListOrdersResponse{}
	for _, order := range orders {
		listOrdersResp = append(listOrdersResp, &ListOrderEntry{
			ID:       order.ID,
			Distance: order.Distance,
			Status:   order.Status,
		})
	}
	return listOrdersResp
}
