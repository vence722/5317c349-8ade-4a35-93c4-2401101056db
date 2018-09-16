package models

type CreateOrderRequest struct {
	Origin      []string `json:"origin"`
	Destination []string `json:"destination"`
}
