package validators

import (
	"5317c349-8ade-4a35-93c4-2401101056db/models"
	"errors"
)

var ValidStatus = []string{"taken"}

func ValidateTakeOrderRequest(req *models.TakeOrderRequst) error {
	// verify status
	valid := false
	for _, status := range ValidStatus {
		if status == req.Status {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("status must be a valid one (taken)")
	}
	return nil
}
