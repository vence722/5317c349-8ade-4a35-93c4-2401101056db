package validators

import (
	"5317c349-8ade-4a35-93c4-2401101056db/models"
	"errors"
)

func ValidateListOrdersRequest(req *models.ListOrdersRequest) error {
	if req.Page <= 0 {
		return errors.New("page must be integer > 0")
	}
	if req.Limit <= 0 {
		return errors.New("limit must be integer > 0")
	}
	return nil
}
