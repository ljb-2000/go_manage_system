package controllers

import (
	"github.com/astaxie/beego"
)


type IndexController struct {
	beego.Controller
}

// @router / [get]
func (this *IndexController) Index() {
	this.TplName = "index.html"
	this.Render()
}
