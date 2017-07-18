package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"strings"
	"net/http"
)

// 用户和用户权限相关的控制器

type UserController struct {
	beego.Controller
}

// 登录信息
type UserInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}


// @router /authenticate [post]
func (this *UserController) Authenticate() {
	// 验证用户身份

	// 返回数据
	result := map[string]interface{}{"data": map[string]interface{}{}, "msg": "响应成功", "code": 200}

	// 1. 获取请求的 账号+密码
	var info UserInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &info)

	//  如果是管理员账号， 管理员登录 管理员账户信息不放在配置文件时，删除这个判断
	if a := beego.AppConfig.String("admin_account"); strings.Compare(a, info.Account) == 0 {
		this.adminLogin(info)
	}

	// 员工登录
	// 2. 数据库中验证账号密码
	flag := strings.Compare("wss", info.Account) == 0

	if !flag {
		// 登录失败
		result["msg"] = "账号或密码错误"
		result["code"] = 10003
		this.Data["json"] = result
		this.ServeJSON()
	}

	// login
	// 账号密码正确， session记录 传递给前端跳转地址
	this.SetSession("user", info.Account)
	result["data"] = map[string]interface{}{"to": "http://localhost:8080/oa/home"}
	result["msg"] = "成功登录"
	result["code"] = http.StatusOK
	this.Data["json"] = result
	this.ServeJSON()
}

// 管理员登录
func (this *UserController) adminLogin(info UserInfo) {
	// 返回数据
	result := map[string]interface{}{"data": map[string]interface{}{}, "msg": "响应成功", "code": 200}

	// 1. 查询管理员的账号密码
	admin_account := beego.AppConfig.String("admin_account")
	admin_password := beego.AppConfig.String("admin_pwd")

	// 2. 验证账号密码
	if strings.Compare(admin_password, info.Password) != 0 {
		// 登录失败
		result["msg"] = "密码错误"
		result["code"] = 10003
		this.Data["json"] = result
		this.ServeJSON()
	}

	// 3. 登录成功
	this.SetSession("user", admin_account)
	result["msg"] = "成功登录"
	result["code"] = http.StatusOK
	this.Data["json"] = result
	this.ServeJSON()

}


// @router /authenticate [post]
func (this *UserController) FindAccess() {
	// 获得用户权限

	// 返回数据
	result := map[string]interface{}{"data": map[string]interface{}{}, "msg": "响应成功", "code": 200}

	// 1. 获得account参数
	account := this.GetString("account")

	if account == "" {
		// 缺少account参数 报错
		result["code"] = 10003
		result["msg"] = "查询失败，缺少账号信息"
		this.Data["json"] = result
		this.ServeJSON()
	}

	result["code"] = http.StatusOK
	result["msg"] = "查询成功"
	this.Data["json"] = result
	this.ServeJSON()
}

// @router /logout [get]
func (this *UserController) Logout() {

	if user := this.GetSession("user"); user != nil {
		//this.DelSession("user")
		this.DestroySession()
		this.Redirect("/oa/login", 302)
	}

}

