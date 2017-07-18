package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

// @router /login [get]
func (this *LoginController) Login() {
	// 用户登录
	this.TplName = "login.html"
	this.Render()
}
