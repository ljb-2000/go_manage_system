package controllers

import (
	"git.fenggese.com/go_manage_system/constant"
	"git.fenggese.com/go_manage_system/logic"
	"git.fenggese.com/go_manage_system/models"
	"github.com/astaxie/beego"
	"net/http"
)

// 用户控制器
type UserController struct {
	beego.Controller
}

// @Title 用户登录
// @router /login [post]
func (this *UserController) Login() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 用户信息
	user := models.UserInfo{}
	if err := this.ParseForm(&user); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if user.Account == "" || user.Password == "" {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 2. 如果是管理员账号， 管理员登录 管理员账户信息不放在配置文件时，删除这个判断
	if admin_account := beego.AppConfig.String("admin_account"); admin_account == user.Account {
		ok := logic.CheckAdmin(&user)
		if !ok {
			// 管理员密码错误
			this.Ctx.Output.JSON(constant.RESP_CODE_ADMIN_LOGIN_ERROR, true, false)
			return
		}
		// 管理员成功登录
		result.Msg = "管理员成功登录"
		this.Ctx.Output.JSON(result, true, false)
		return
	}

	// 3. 用户登录
	// 验证账号密码
	ok, err := logic.LoginLogic(&user)
	if err != nil {
		// 登录出错，数据库查询错误
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 登录失败
	if !ok {
		this.Ctx.Output.JSON(constant.RESP_CODE_LOGIN_ERROR, true, false)
		return
	}

	// 用户成功登录
	result.Msg = "用户成功登录"
	this.Ctx.Output.JSON(result, true, false)
}
