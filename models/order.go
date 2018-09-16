package models

import "time"

// The order model
type Order struct {
	ID                   string    `orm:"column(id);pk"`
	OriginLatitude       float64   `orm:"column(origin_latitude)"`
	OriginLongitude      float64   `orm:"column(origin_longitude)"`
	DestinationLatitude  float64   `orm:"column(destination_latitude)"`
	DestinationLongitude float64   `orm:"column(destination_longitude)"`
	Distance             float64   `orm:"column(distance)"`
	Status               string    `orm:"column(status)"` // enum of UNASSIGN, TAKEN
	CreatedAt            time.Time `orm:"column(created_at)"`
	UpdatedAt            time.Time `orm:"column(updated_at)"`
}
