package controllers

import (
	"github.com/astaxie/beego"
	"git.gumpcome.com/gumpoa/models"
	"git.gumpcome.com/gumpoa/logic"
	"net/http"
	"git.gumpcome.com/gumpoa/constant"
	"encoding/json"
)

type AccessController struct {
	beego.Controller
}


// @Title 添加权限
// @Description
// @reveive 权限名: name, 权限码: code
// @router /add [post]
func (this *AccessController) AccessAdd() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 权限信息
	access := models.Access{}
	if err := this.ParseForm(&access); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if access.Code <= 0 || access.Name == "" {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 添加权限
	isAdd, _, err := logic.AddAccessLogic(&access)
	if err != nil {
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 处理结果
	if !isAdd {
		// 添加失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ACCESS_ADD_ERROR, true, false)
		return
	}
	// 添加成功
	result.Msg = "权限添加成功"
	this.Ctx.Output.JSON(result, true, false)
}


// @Title 删除权限
// @receive 权限: access
// @router /delete [delete]
func (this *AccessController) AccessDelete() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 权限信息
	access := models.Access{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &access); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if access.Code <= 0 {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 删除权限
	isDelete, err := logic.DeleteAccessLogic(&access)
	if err != nil {
		// 删除出错
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 处理结果
	if !isDelete {
		// 删除失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ACCESS_DELETE_ERROR, true, false)
		return
	}
	// 删除成功
	result.Msg = "权限删除成功"
	this.Ctx.Output.JSON(result, true, false)
}

//
//// @Title 根据账号查找权限列表
//// @receive 账号: account
//// @router /get_menus [get]
//func (this *AccessController) CreateMenus() {
//	// 声明响应结构体
//	result := models.CommonWithDataResp{Code: http.StatusOK}
//
//	// 1. 获取并解析请求的 权限信息
//	account := this.GetString("account")
//
//	// 2. 检查参数
//	if account == "" {
//		// 参数不合法
//		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
//		return
//	}
//
//	// 3. 生成菜单
//	menus, err := logic.CreateMenusLogic(&account)
//	if  err != nil {
//		// 生成权限菜单失败
//		this.Ctx.Output.JSON(constant.RESP_CODE_ACCESS_MENUS_ERROR, true, false)
//		return
//	}
//	// 生成权限菜单成功
//	result.Msg = "菜单查询成功"
//	result.Data = map[string]interface{}{"menus": menus}
//	this.Ctx.Output.JSON(result, true, false)
//}

// @Title 权限列表
// @receive 页数: page_number 每页结果数: page_size
// @router /list [post]
func (this *AccessController) AccessList() {
	// 声明响应结构体
	result := models.CommonWithDataResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 分页信息
	filter := models.PageAccessFilter{}
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
	page, err := logic.PageAccessLogic(&filter)
	if err != nil {
		// 权限分页查找失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ACCESS_PAGE_ERROR, true, false)
		return
	}

	// 查询权限列表成功
	result.Msg = "权限列表查询成功"
	result.Data = map[string]interface{}{"page": page}
	this.Ctx.Output.JSON(result, true, false)
}