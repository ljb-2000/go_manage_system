package logic

import (
	"git.gumpcome.com/gumpoa/models"
	"git.gumpcome.com/gumpoa/dao"
	"git.gumpcome.com/gumpoa/constant"
	"git.gumpcome.com/go_kit/dbkit"
)

// @Title 添加权限
func AddAccessLogic(access *models.Access) (bool, int64, error) {

	// 1. 检查参数
	if access.Code <= 0 || access.Name == "" {
		// 参数不合法
		return false, 0, constant.RESP_CODE_PARAMS_ERROR
	}

	// 2. 准备数据
	params := map[string]interface{}{
		"code": access.Code,
		"name": access.Name,
	}

	// 3. 添加权限
	return dao.AddAccessDao(&params)
}

// @Title 删除权限
func DeleteAccessLogic(access *models.Access) (bool, error) {

	// 1. 检查参数
	if access.Code <= 0 {
		// 参数不合法
		return false, constant.RESP_CODE_PARAMS_ERROR
	}

	// 2. 准备数据
	params := map[string]interface{}{
		"code": access.Code,
	}

	// 3. 删除权限
	return dao.DeleteAccessDao(&params)
}

// @Title 分页查找权限
func PageAccessLogic(filter *models.PageAccessFilter) (dbkit.Page, error)  {
	return dao.PageAccessDao(filter)
}

