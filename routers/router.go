package routers

import (
	"git.fenggese.com/go_manage_system/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	// 跨域访问请求，开发时候要用，部署时不需要
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	// 主页
	beego.Router("/", &controllers.IndexController{}, "Get:Index")

	ns := beego.NewNamespace("/api.fenggese.com/manage",
		// api
		beego.NSNamespace("/v1",
			// /api.fenggese.com/manage/v1/user
			beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
			// /api.fenggese.com/manage/v1/access
			beego.NSNamespace("/access", beego.NSInclude(&controllers.AccessController{})),
			// /api.fenggese.com/manage/v1/role
			beego.NSNamespace("/role", beego.NSInclude(&controllers.RoleController{})),
			// /api.fenggese.com/manage/v1/account
			beego.NSNamespace("/account", beego.NSInclude(&controllers.AccountController{})),
		),
	)
	beego.AddNamespace(ns)
}
