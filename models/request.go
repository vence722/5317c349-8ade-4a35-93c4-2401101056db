package models

type CreateOrderRequest struct {
	Origin      []string `json:"origin"`
	Destination []string `json:"destination"`
}

type TakeOrderRequst struct {
	OrderID string
	Status  string `json:"status"`
}
