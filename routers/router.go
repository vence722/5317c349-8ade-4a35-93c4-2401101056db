package routers

import (
	"5317c349-8ade-4a35-93c4-2401101056db/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/order", &controllers.CreateOrderController{})
	beego.Router("/order/:id", &controllers.TakeOrderController{})
	beego.Router("/orders", &controllers.ListOrdersController{})
}
