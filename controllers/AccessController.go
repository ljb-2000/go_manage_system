package controllers

import (
	"github.com/astaxie/beego"
	"git.gumpcome.com/gumpoa/models"
	"git.gumpcome.com/gumpoa/logic"
	"net/http"
	"git.gumpcome.com/gumpoa/constant"
)

type AccessController struct {
	beego.Controller
}


// @Title 添加权限
// @Description
//	接收： 权限名：code, 权限码：code
// @router /add [post]
func (this *AccessController) AccessAdd() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK, Msg: ""}

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
// @Description
//	接收： 权限码：code
// @router /delete [delete]
func (this *AccessController) AccessDelete() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK, Msg: ""}

	// 1. 获取并解析请求的 权限信息
	access := models.Access{}
	if err := this.ParseForm(&access); err != nil {
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
	return
}


// @Title 根据账号查找权限列表
// @router /get_menus [get]
func (this *AccessController) CreateMenus() {
	// 声明响应结构体
	result := models.CommonWithDataResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 权限信息
	account := this.GetString("account")

	// 2. 检查参数
	if account == "" {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 生成菜单
	menus, ok := logic.CreateMenusLogic(account)
	if !ok {
		// 生成权限菜单失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ACCESS_LIST_ERROR, true, false)
		return
	}
	// 生成权限菜单成功
	result.Data = map[string]interface{}{"menus": menus}
	this.Ctx.Output.JSON(result, true, false)
	return

}