package services

import (
	"5317c349-8ade-4a35-93c4-2401101056db/models"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"github.com/vence722/convert"
	"googlemaps.github.io/maps"
)

func CreateOrder(req *models.CreateOrderRequest) (*models.Order, error) {
	// call google api to get distance
	distance, err := getDistance(req.Origin, req.Destination)
	if err != nil {
		return nil, err
	}

	// insert order to MySQL
	o := orm.NewOrm()
	order := &models.Order{
		ID:                   uuid.Must(uuid.NewV4()).String(),
		OriginLatitude:       convert.Str2Float64(req.Origin[0]),
		OriginLongitude:      convert.Str2Float64(req.Origin[1]),
		DestinationLatitude:  convert.Str2Float64(req.Destination[0]),
		DestinationLongitude: convert.Str2Float64(req.Destination[1]),
		Distance:             distance,
		Status:               "UNASSIGN",
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	_, err = o.Insert(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func TakeOrder(req *models.TakeOrderRequst) (*models.Order, error) {
	o := orm.NewOrm()

	// start transaction
	o.Begin()

	// get order by orderID
	order := &models.Order{
		ID: req.OrderID,
	}
	err := o.Read(order)
	if err != nil {
		o.Rollback()
		return nil, errors.New("ORDER_NOT_FOUND")
	}

	// update order status to TAKEN if it's not taken
	if order.Status == "TAKEN" {
		o.Rollback()
		return nil, errors.New("ORDER_ALREADY_BEEN_TAKEN")
	}
	// change status to TAKEN
	order.Status = "TAKEN"
	// update UpdatedAt
	order.UpdatedAt = time.Now()
	_, err = o.Update(order, "Status")
	if err != nil {
		o.Rollback()
		return nil, err
	}

	// commit transation
	o.Commit()

	return order, nil
}

func ListOrders(req *models.ListOrdersRequest) ([]*models.Order, error) {
	o := orm.NewOrm()

	// calculate from, pageSize
	from := (req.Page - 1) * req.Limit
	pageSize := req.Limit

	// get orders from DB
	var orders []*models.Order
	_, err := o.Raw("SELECT * FROM `order` LIMIT ?,?", from, pageSize).QueryRows(&orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func getDistance(origin []string, destination []string) (float64, error) {
	r := &maps.DistanceMatrixRequest{
		Origins:      []string{fmt.Sprintf("%s,%s", origin[0], origin[1])},
		Destinations: []string{fmt.Sprintf("%s,%s", destination[0], destination[1])},
	}
	resp, err := googleMapsClient.DistanceMatrix(context.Background(), r)
	if err != nil {
		return 0, err
	}
	distance := resp.Rows[0].Elements[0].Distance.Meters
	return float64(distance), nil
}
