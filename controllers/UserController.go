package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"git.gumpcome.com/gumpoa/logic"
	"git.gumpcome.com/gumpoa/models"
)

// 用户和用户权限相关的控制器
type UserController struct {
	beego.Controller
}



// @Title 用户登出
// @router /logout [get]
func (this *UserController) Logout() {
	if user := this.GetSession("user"); user != nil {
		this.DestroySession()
	}
}

// @Title 用户是否已经登录
// @router /haslogined [get]
func (this *UserController) HasLogined() {

	// 返回数据
	result := models.NewResult()

	// 用户没有登录
	if user := this.GetSession("user"); user == nil {
		result["msg"] = "用户没有登录"
		result["code"] = models.BusinessWrong
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	// 用户已经登录
	result["msg"] = "用户已经登录"
	result["code"] = 200
	this.Data["json"] = result
	this.ServeJSON()
	return
}

// @Title 用户登录
// @router /login [post]
func (this *UserController) Login() {

	// 1. 获取请求的 账号+密码
	var info models.UserInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	//fmt.Println(string(this.Ctx.Input.RequestBody))

	//  如果是管理员账号， 管理员登录 管理员账户信息不放在配置文件时，删除这个判断
	if admin_account := beego.AppConfig.String("admin_account"); admin_account == info.Account {
		this.Data["json"], _ = logic.CheckAdmin(info)
		this.ServeJSON()
		return
	}

	// 员工登录
	// 2. 验证账号密码
	var ok bool
	this.Data["json"], ok = logic.CheckUser(info)
	if ok {
		// 账号密码正确, session保存用户状态
		this.SetSession("user", info.Account)
	}

	this.ServeJSON()
	return
}
