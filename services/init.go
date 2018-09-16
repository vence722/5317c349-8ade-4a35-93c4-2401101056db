package services

import (
	"github.com/astaxie/beego/orm"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"5317c349-8ade-4a35-93c4-2401101056db/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:peggy2015@/orders?charset=utf8")
	orm.RegisterModel(new(models.Order))
}
