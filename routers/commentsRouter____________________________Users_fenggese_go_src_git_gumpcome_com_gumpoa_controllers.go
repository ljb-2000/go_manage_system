package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:HomeController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Home",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:LoginController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"],
		beego.ControllerComments{
			Method: "Authenticate",
			Router: `/authenticate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"],
		beego.ControllerComments{
			Method: "FindAccess",
			Router: `/authenticate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
