package logic

import (
	"git.fenggese.com/go_kit/dbkit"
	"git.fenggese.com/go_manage_system/constant"
	"git.fenggese.com/go_manage_system/dao"
	"git.fenggese.com/go_manage_system/models"
)

// @Title 添加角色
func AddRoleLogic(role *models.Role) (bool, int64, error) {

	// 1. 检查参数
	if role.Name == "" {
		// 参数不合法
		return false, 0, constant.RESP_CODE_PARAMS_VALUE_ERROR
	}

	// 2. 准备数据
	params := map[string]interface{}{
		"name": role.Name,
	}

	// 3. 添加角色
	return dao.AddRoleDao(&params)
}

// @Title 绑定权限
func BindAccessLogic(roleAccess *models.RoleAccess) (bool, int64, error) {

	// 1. 检查参数
	if roleAccess.RoleID == 0 || roleAccess.AccessID == 0 {
		// 参数不合法
		return false, 0, constant.RESP_CODE_PARAMS_VALUE_ERROR
	}

	// 2. 准备数据
	params := map[string]interface{}{
		"role_id":   roleAccess.RoleID,
		"access_id": roleAccess.AccessID,
	}

	// 3. 绑定权限
	return dao.BindAccessDao(&params)
}

// @Title 解绑权限
func UnbindAccessLogic(roleAccess *models.RoleAccess) (bool, error) {

	// 1. 检查参数
	if roleAccess.RoleID == 0 && roleAccess.AccessID == 0 {
		// 参数不合法
		return false, constant.RESP_CODE_PARAMS_VALUE_ERROR
	}

	// 2. 准备数据
	params := map[string]interface{}{
		"role_id":   roleAccess.RoleID,
		"access_id": roleAccess.AccessID,
	}

	// 3. 解除权限
	return dao.UnbindAccessDao(&params)
}

// @Title 分页查找权限
func PageRoleLogic(filter *models.PageRoleFilter) (dbkit.Page, error) {
	// 1. 检查参数
	if filter.PageNumber < 1 || filter.PageSize < 0 {
		// 参数不合法
		return nil, constant.RESP_CODE_PARAMS_VALUE_ERROR
	}

	// 2. 分页查找
	return dao.PageRoleDao(filter)
}

// @Title 查询角色的权限
func AccessQueryByRoleLogic(role_id int64) ([]map[string]interface{}, error) {
	// 1. 检查参数
	if role_id <= 0 {
		// 参数不合法
		return nil, constant.RESP_CODE_PARAMS_VALUE_ERROR
	}

	// 2. 查询
	return dao.AccessQueryByRoleDao(role_id)
}

// @Title 删除角色
func DeleteRoleLogic(role_id int) (bool, error) {
	// 1. 检查参数
	if role_id <= 0 {
		// 参数不合法
		return nil, constant.RESP_CODE_PARAMS_VALUE_ERROR
	}

	// 2. 删除
	return dao.DeleteRoleDao(role_id)
}
