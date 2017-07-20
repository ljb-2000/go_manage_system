package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:AccessController"],
		beego.ControllerComments{
			Method: "FindList",
			Router: `/find_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:IndexController"] = append(beego.GlobalControllerRouter["git.gumpcome.com/gumpoa/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
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
