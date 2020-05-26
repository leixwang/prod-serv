package main

import (
	_ "prod-serv/models"
	_ "prod-serv/routers"
	_ "prod-serv/utils"

	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	// beego.BConfig.WebConfig.Session.SessionOn = true

	// beego.InsertFilter("/api/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run("0.0.0.0:9000")

}
