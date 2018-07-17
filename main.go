package main

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	_ "github.com/udistrital/administrativa_crud_api/routers"
	"github.com/udistrital/utils_oas/apiStatusLib"
)

func init() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC
	q := "postgres://" + beego.AppConfig.String("PGuser") + ":" + beego.AppConfig.String("PGpass") + "@" + beego.AppConfig.String("PGurls") + "/" + beego.AppConfig.String("PGdb") + "?sslmode=disable&search_path=" + beego.AppConfig.String("PGschemas") + "&timezone=UTC"
	//fmt.Println(q)
	orm.RegisterDataBase("default", "postgres", q)
}

func main() {
	orm.Debug = true

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Custom JSON error pages
	beego.ErrorHandler("403", forgivenJsonPage)
	beego.ErrorHandler("404", notFoundJsonPage)

	logs.SetLogger(logs.AdapterFile, `{"filename":"/var/log/beego/administrativa_crud_api.log"}`)

	apistatus.Init()
	beego.Run()
}
