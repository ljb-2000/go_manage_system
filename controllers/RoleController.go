package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"git.gumpcome.com/gumpoa/models"
	"git.gumpcome.com/gumpoa/constant"
	"git.gumpcome.com/gumpoa/logic"
	"encoding/json"
)

type RoleController struct {
	beego.Controller
}


// @Title 添加角色
// @reveive 角色名: name
// @router /add [post]
func (this *RoleController) RoleAdd() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 角色信息
	role := models.Role{}
	if err := this.ParseForm(&role); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if role.Name == "" {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 添加角色
	isAdd, _, err := logic.AddRoleLogic(&role)
	if err != nil {
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 处理结果
	if !isAdd {
		// 添加失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ROLE_ADD_ERROR, true, false)
		return
	}
	// 添加成功
	result.Msg = "角色添加成功"
	this.Ctx.Output.JSON(result, true, false)
}


// @Title 为角色绑定权限
// @reveive 角色ID: role_id 权限ID: access_id
// @router /bind [post]
func (this *RoleController) RoleAccessBind() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 角色ID和权限ID
	roleAccess := models.RoleAccess{}
	if err := this.ParseForm(&roleAccess); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if roleAccess.RoleID == 0 || roleAccess.AccessID == 0 {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 绑定权限
	isAdd, _, err := logic.BindAccessLogic(&roleAccess)
	if err != nil {
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 处理结果
	if !isAdd {
		// 绑定权限失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ROLE_ACCESS_BIND_ERROR, true, false)
		return
	}
	// 绑定成功
	result.Msg = "为角色绑定权限成功"
	this.Ctx.Output.JSON(result, true, false)
}


// @Title 为角色解绑权限
// @reveive （ 角色ID: role_id 权限ID: access_id ）或（ 角色权限表ID: id ）
// @router /unbind [post]
func (this *RoleController) RoleAccessUnbind() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的
	roleAccess := models.RoleAccess{}
	if err := this.ParseForm(&roleAccess); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if roleAccess.RoleID == 0 && roleAccess.AccessID == 0 {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 解绑权限
	isUnbind, err := logic.UnbindAccessLogic(&roleAccess)
	if err != nil {
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 处理结果
	if !isUnbind {
		// 解除权限失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ROLE_ACCESS_UNBIND_ERROR, true, false)
		return
	}
	// 解绑成功
	result.Msg = "为角色解除权限成功"
	this.Ctx.Output.JSON(result, true, false)
}



// @Title 角色列表
// @receive 页数: page_number 每页结果数: page_size
// @router /list [post]
func (this *RoleController) RoleList() {
	// 声明响应结构体
	result := models.CommonWithDataResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 分页信息
	filter := models.PageRoleFilter{}
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
	page, err := logic.PageRoleLogic(&filter)
	if err != nil {
		// 角色分页查找失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ROLE_PAGE_ERROR, true, false)
		return
	}

	// 查询角色列表成功
	result.Msg = "角色列表查询成功"
	result.Data = map[string]interface{}{"page": page}
	this.Ctx.Output.JSON(result, true, false)
}



// @Title 查询角色的权限
// @reveive 角色ID: role_id
// @router /query_access [get]
func (this *RoleController) AccessQuery() {
	// 声明响应结构体
	result := models.CommonWithDataResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 角色信息
	role_id, err := this.GetInt64("role_id")
	if err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if role_id == 0 {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 查询权限
	accesses, err := logic.AccessQueryLogic(role_id)
	if err != nil {
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 添加成功
	result.Msg = "角色权限查询成功"
	result.Data = map[string]interface{}{"accesses": accesses}
	this.Ctx.Output.JSON(result, true, false)
}



// @Title 删除权限
// @receive 角色: role
// @router /delete [delete]
func (this *RoleController) RoleDelete() {
	// 声明响应结构体
	result := models.CommonResp{Code: http.StatusOK}

	// 1. 获取并解析请求的 角色信息
	role := models.Role{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &role); err != nil {
		// 参数错误
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_ERROR, true, false)
		return
	}

	// 2. 检查参数
	if role.ID == 0 {
		// 参数不合法
		this.Ctx.Output.JSON(constant.RESP_CODE_PARAMS_VALUE_ERROR, true, false)
		return
	}

	// 3. 删除权限
	isDelete, err := logic.DeleteRoleLogic(role.ID)
	if err != nil {
		// 删除出错
		this.Ctx.Output.JSON(err, true, false)
		return
	}

	// 4. 处理结果
	if !isDelete {
		// 删除失败
		this.Ctx.Output.JSON(constant.RESP_CODE_ROLE_DELETE_ERROR, true, false)
		return
	}
	// 删除成功
	result.Msg = "角色删除成功"
	this.Ctx.Output.JSON(result, true, false)
}