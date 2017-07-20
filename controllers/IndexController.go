package controllers

import (
	"github.com/astaxie/beego"
)

// 存放需要单点登入的url
type IndexController struct {
	beego.Controller
}

// @router / [get]
func (this *IndexController) Index() {
	this.TplName = "index.html"
	this.Render()
}
