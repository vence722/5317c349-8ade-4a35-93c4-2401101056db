package validators

import (
	"5317c349-8ade-4a35-93c4-2401101056db/models"
	"errors"

	"gopkg.in/asaskevich/govalidator.v4"
)

func ValidateCreateOrderRequest(req *models.CreateOrderRequest) error {
	origin := req.Origin
	destination := req.Destination
	// validate array length
	if len(origin) != 2 {
		return errors.New("origin must have two elements of float type")
	}
	if len(destination) != 2 {
		return errors.New("destination must have two elements of float type")
	}
	// validate data type
	if !govalidator.IsFloat(origin[0]) || !govalidator.IsFloat(origin[1]) {
		return errors.New("origin latitude and longitude must be float")
	}
	if !govalidator.IsFloat(destination[0]) || !govalidator.IsFloat(destination[1]) {
		return errors.New("destination latitude and longitude must be float")
	}
	return nil
}
