package routers

import (
	"git.gumpcome.com/gumpoa/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 主页
	beego.Router("/", &controllers.HomeController{}, "Get:Home")

	ns := beego.NewNamespace("/oa",
		// web
		beego.NSInclude(&controllers.LoginController{}),

		// api
		beego.NSNamespace("/v1",
			beego.NSNamespace("/api",
				beego.NSNamespace("/user",
					// /oa/v1/api/user
					beego.NSInclude(&controllers.UserController{}),
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
