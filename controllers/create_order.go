package controllers

import (
	"5317c349-8ade-4a35-93c4-2401101056db/controllers/validators"
	"5317c349-8ade-4a35-93c4-2401101056db/models"
	"5317c349-8ade-4a35-93c4-2401101056db/services"
	"encoding/json"

	"github.com/astaxie/beego"
)

type CreateOrderController struct {
	beego.Controller
}

func (this *CreateOrderController) Post() {
	// get request params
	req := &models.CreateOrderRequest{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, req)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = models.NewErrorResponse(err)
		this.ServeJSON()
		return
	}

	// validate request params
	err = validators.ValidateCreateOrderRequest(req)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = models.NewErrorResponse(err)
		this.ServeJSON()
		return
	}

	// call create order service
	order, err := services.CreateOrder(req)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = models.NewErrorResponse(err)
		this.ServeJSON()
		return
	}

	// return success response
	this.Data["json"] = models.NewCreateOrderResponse(order)
	this.ServeJSON()
}
