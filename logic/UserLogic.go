package logic

import (
	"github.com/astaxie/beego"
	"git.gumpcome.com/gumpoa/models"
)

// @Title 验证管理员账号密码
func CheckAdmin(info *models.UserInfo) bool {
	// 1. 查询管理员的账号密码
	admin_account := beego.AppConfig.String("admin_account")
	admin_password := beego.AppConfig.String("admin_pwd")

	// 2. 验证账号密码
	if info.Account != admin_account || info.Password != admin_password {
		// 账号密码不匹配
		return false
	}
	// 成功登录
	return true
}

// @Title 验证用户账号密码
func CheckUser(info *models.UserInfo) bool {
	// TODO 待完善
	if info.Account != "wss" {
		return false
	}
	return true
}