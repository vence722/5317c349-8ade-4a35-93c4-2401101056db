package services

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"googlemaps.github.io/maps"

	"5317c349-8ade-4a35-93c4-2401101056db/models"
)

var googleMapsClient *maps.Client

func init() {
	// init mongo db
	user := beego.AppConfig.String("mysql_user")
	pass := beego.AppConfig.String("mysql_pass")
	db := beego.AppConfig.String("mysql_db")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", user, pass, db))
	orm.RegisterModel(new(models.Order))

	// init google map api
	apiKey := beego.AppConfig.String("google_map_api_key")
	googleMapsClient, _ = maps.NewClient(maps.WithAPIKey(apiKey))
}
