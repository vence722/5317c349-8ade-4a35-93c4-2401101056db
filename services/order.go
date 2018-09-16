package services

import (
	"5317c349-8ade-4a35-93c4-2401101056db/models"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"github.com/vence722/convert"
)

func CreateOrder(req *models.CreateOrderRequest) (*models.Order, error) {
	o := orm.NewOrm()
	order := &models.Order{
		ID:                   uuid.Must(uuid.NewV4()).String(),
		OriginLatitude:       convert.Str2Float64(req.Origin[0]),
		OriginLongitude:      convert.Str2Float64(req.Origin[1]),
		DestinationLatitude:  convert.Str2Float64(req.Destination[0]),
		DestinationLongitude: convert.Str2Float64(req.Destination[1]),
		Distance:             0.0,
		Status:               "UNASSIGN",
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	_, err := o.Insert(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
