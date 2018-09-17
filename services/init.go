package services

import (
	"fmt"
	"time"

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
	env := beego.AppConfig.String("runmode")
	host := beego.AppConfig.String(env + "::mysql_host")
	user := beego.AppConfig.String(env + "::mysql_user")
	pass := beego.AppConfig.String(env + "::mysql_pass")
	db := beego.AppConfig.String(env + "::mysql_db")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	for {
		err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, pass, host, db))
		if err == nil {
			break
		}
		fmt.Println("MySQL is not ready, sleep 10s...")
		time.Sleep(10 * time.Second)
	}
	orm.RegisterModel(new(models.Order))

	// init google map api
	apiKey := beego.AppConfig.String("google_map_api_key")
	googleMapsClient, _ = maps.NewClient(maps.WithAPIKey(apiKey))
}
