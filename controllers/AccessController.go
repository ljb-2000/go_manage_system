package controllers

import (
	"github.com/astaxie/beego"
	"git.gumpcome.com/gumpoa/models"
	"git.gumpcome.com/gumpoa/logic"
)

type AccessController struct {
	beego.Controller
}


// @Title 根据账号查找权限列表
// @router /find_list [get]
func (this *AccessController) FindList() {
	// 返回数据
	result := models.NewResult()

	account := this.GetString("account")

	// 是否有用户账号
	if account == "" {
		result["msg"] = "权限查询失败, 缺少用户账号"
		result["code"] = models.BusinessWrong
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	tmp, ok := logic.CreateMenus(account)

	// 权限是否查出
	if !ok {
		result["msg"] = "权限查询失败"
		result["code"] = models.BusinessWrong
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	// 查出权限
	result["msg"] = "权限查询成功"
	result["code"] = 200
	result["data"] = tmp
	this.Data["json"] = result
	this.ServeJSON()
	return
}