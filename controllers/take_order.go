package controllers

import (
	"5317c349-8ade-4a35-93c4-2401101056db/controllers/validators"
	"5317c349-8ade-4a35-93c4-2401101056db/models"
	"5317c349-8ade-4a35-93c4-2401101056db/services"
	"encoding/json"

	"github.com/astaxie/beego"
)

type TakeOrderController struct {
	beego.Controller
}

func (this *TakeOrderController) Put() {
	// get request params
	req := &models.TakeOrderRequst{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, req)
	if err != nil {
		this.Data["json"] = models.NewErrorResponse(err)
		this.ServeJSON()
		return
	}
	// set order ID
	orderID := this.GetString(":id")
	req.OrderID = orderID

	// validate request params
	err = validators.ValidateTakeOrderRequest(req)
	if err != nil {
		this.Data["json"] = models.NewErrorResponse(err)
		this.ServeJSON()
		return
	}

	// call take order service
	_, err = services.TakeOrder(req)
	if err != nil {
		this.Data["json"] = models.NewErrorResponse(err)
		this.ServeJSON()
		return
	}

	// return success response
	this.Data["json"] = models.NewTakeOrderResponse()
	this.ServeJSON()
}
