package logic

import (
	"github.com/astaxie/beego"
	"net/http"
	"git.gumpcome.com/gumpoa/models"
)

// @Title 验证管理员账号密码
func CheckAdmin(info models.UserInfo) (models.Result, bool) {
	// 返回数据
	result := models.NewResult()

	// 1. 查询管理员的账号密码
	admin_account := beego.AppConfig.String("admin_account")
	admin_password := beego.AppConfig.String("admin_pwd")

	if info.Account != admin_account || info.Password != admin_password {
		// 账号或密码错误
		result["msg"] = "管理员密码错误"
		result["code"] = models.BusinessWrong
		return result, false
	}

	// 账号密码正确
	result["msg"] = "管理员登录成功"
	result["code"] = http.StatusOK
	return result, true

}

// @Title 验证用户账号密码
func CheckUser(info models.UserInfo) (models.Result, bool) {
	// 返回数据
	result := models.NewResult()

	// 1. 查询数据库
	account := "wss"

	if info.Account != account {
		// 账号或密码错误
		result["msg"] = "账号或密码错误"
		result["code"] = models.BusinessWrong
		return result, false
	}

	// 账号密码正确
	result["msg"] = "用户登录成功"
	result["code"] = http.StatusOK
	return result, true
}