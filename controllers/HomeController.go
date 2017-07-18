package controllers

import (
	"github.com/astaxie/beego"
)

// 存放需要单点登入的url

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Prepare() {
	// 单点登入
	if user := this.GetSession("user"); user == nil {
		this.Redirect("/oa/login", 302)
	}

}


// 左侧菜单
type Menus struct {
	Index int
	Name  string
	Url   string
	Clazz string
}

// @router / [get]
func (this *HomeController) Home() {

	this.Data["user"] = this.GetSession("user")

	// 测试用数据
	menus := []Menus{
		{1, "权限管理", "#", "el-icon-document"},
		{2, "角色管理", "#", "el-icon-document"},
		{3, "账号管理", "#", "el-icon-document"},
		{4, "个人设置", "#", "el-icon-setting"},
	}

	this.Data["menus"] = menus
	this.TplName = "home.html"
	this.Render()
}
