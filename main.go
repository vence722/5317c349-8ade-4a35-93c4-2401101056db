package main

import (
	_ "5317c349-8ade-4a35-93c4-2401101056db/middlewares"
	_ "5317c349-8ade-4a35-93c4-2401101056db/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
