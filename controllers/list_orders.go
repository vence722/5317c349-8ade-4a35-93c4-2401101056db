package controllers

import (
	"5317c349-8ade-4a35-93c4-2401101056db/controllers/validators"
	"5317c349-8ade-4a35-93c4-2401101056db/models"
	"5317c349-8ade-4a35-93c4-2401101056db/services"
	"errors"

	"github.com/astaxie/beego"
)

type ListOrdersController struct {
	beego.Controller
}

func (this *ListOrdersController) Get() {
	// get request params
	req := &models.ListOrdersRequest{}
	page, err := this.GetInt("page")
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = models.NewErrorResponse(errors.New("page must be integer > 0"))
		this.ServeJSON()
		return
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = models.NewErrorResponse(errors.New("limit must be integer > 0"))
		this.ServeJSON()
		return
	}
	req.Page = page
	req.Limit = limit

	// validate request params
	err = validators.ValidateListOrdersRequest(req)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = models.NewErrorResponse(err)
		this.ServeJSON()
		return
	}

	// call create order service
	orders, err := services.ListOrders(req)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = models.NewErrorResponse(err)
		this.ServeJSON()
		return
	}

	// return success response
	this.Data["json"] = models.NewListOrdersResponse(orders)
	this.ServeJSON()
}
