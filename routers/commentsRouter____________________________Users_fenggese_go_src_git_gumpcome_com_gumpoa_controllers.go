package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"],
		beego.ControllerComments{
			Method: "AccessAdd",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"],
		beego.ControllerComments{
			Method: "AccessDelete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"],
		beego.ControllerComments{
			Method: "CreateMenus",
			Router: `/get_menus`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"],
		beego.ControllerComments{
			Method: "AccessList",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccountController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccountController"],
		beego.ControllerComments{
			Method: "AccountAdd",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccountController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccountController"],
		beego.ControllerComments{
			Method: "AccountDelete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccountController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccountController"],
		beego.ControllerComments{
			Method: "AccountList",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:IndexController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"],
		beego.ControllerComments{
			Method: "RoleAdd",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"],
		beego.ControllerComments{
			Method: "RoleAccessBind",
			Router: `/bind`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"],
		beego.ControllerComments{
			Method: "RoleAccessUnbind",
			Router: `/unbind`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"],
		beego.ControllerComments{
			Method: "RoleList",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"],
		beego.ControllerComments{
			Method: "AccessQuery",
			Router: `/query_access`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:RoleController"],
		beego.ControllerComments{
			Method: "RoleDelete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"],
		beego.ControllerComments{
			Method: "HasLogined",
			Router: `/haslogined`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
