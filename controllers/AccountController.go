package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"git.gumpcome.com/gumpoa/models"
	"git.gumpcome.com/gumpoa/constant"
	"git.gumpcome.com/gumpoa/logic"
	"encoding/json"
)

type AccountController struct {
	beego.Controller
}



// @Title 添加账号
// @reveive 用户: user_name, 账号: login_name 密码: login_pwd 角色ID: role_id
// @router /add [post]
func (this *AccountController) AccountAdd() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 账号信息
	account := models.Account{}
	if err := this.ParseForm(&account); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if account.RoleID == 0 || account.LoginName == "" || account.LoginPwd == "" || account.UserName == "" {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 添加账号
	isAdd, _, err := logic.AddAccountLogic(&account)
	if err != nil {
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 处理结果
	if !isAdd {
		// 添加失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ACCOUNT_ADD_ERROR, true, false)
		return
	}
	// 添加成功
	result.Msg = "账号添加成功"
	this.Ctx.Output.JSON(result, true, false)
}


// @Title 删除账号
// @receive 账号: account
// @router /delete [delete]
func (this *AccountController) AccountDelete() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的  账号信息
	account := models.Account{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &account); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if account.ID == 0 {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 删除账号
	isDelete, err := logic.DeleteAccountLogic(&account)
	if err != nil {
		// 删除出错
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 处理结果
	if !isDelete {
		// 删除失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ACCOUNT_DELETE_ERROR, true, false)
		return
	}
	// 删除成功
	result.Msg = "账号删除成功"
	this.Ctx.Output.JSON(result, true, false)
}

// @Title 账号列表
// @receive 页数: page_number 每页结果数: page_size
// @router /list [post]
func (this *AccountController) AccountList() {
	// 声明响应结构体
	result := models.CommonWithDataResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 分页信息
	filter := models.PageAccountFilter{}
	if err := this.ParseForm(&filter); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if filter.PageNumber < 1 || filter.PageSize < 0 {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 查询权限列表
	page, err := logic.PageAccountLogic(&filter)
	if err != nil {
		// 权限分页查找失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ACCOUNT_PAGE_ERROR, true, false)
		return
	}

	// 查询账号列表成功
	result.Msg = "账号列表查询成功"
	result.Data = map[string]interface{}{"page": page}
	this.Ctx.Output.JSON(result, true, false)
}


// @Title 根据账号查找权限列表
// @receive 账号 login_name
// @router /access_list [post]
func (this *AccountController)AccountAccessListQuery() {
	// 声明响应结构体
	result := models.CommonWithDataResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 账号信息
	account := models.Account{}
	if err := this.ParseForm(&account); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if account.LoginName == "" {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 查询权限
	accesses,  err := logic.AccessQueryByAccountLogic(&account)
	if err != nil {
		// 根据账号查找权限失败
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 查找成功
	result.Msg = "根据账号查找权限列表成功"
	result.Data = map[string]interface{}{"accesses": accesses}
	this.Ctx.Output.JSON(result, true, false)
}



// @Title 根据账号生成菜单
// @receive 账号 login_name
// @router /menus_list [post]
func (this *AccountController) MenusQuery() {
	// 声明响应结构体
	result := models.CommonWithDataResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 账号信息
	account := models.Account{}
	if err := this.ParseForm(&account); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if account.LoginName == "" {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 生成菜单
	menus,  err := logic.CreateMenusLogic(&account)
	if err != nil {
		// 根据账号生成菜单失败
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 查找成功
	result.Msg = "根据账号生成菜单成功"
	result.Data = map[string]interface{}{"menus": menus}
	this.Ctx.Output.JSON(result, true, false)
}
