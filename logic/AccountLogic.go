package logic

import (
	"git.gumpcome.com/gumpoa/models"
	"git.gumpcome.com/gumpoa/constant"
	"git.gumpcome.com/gumpoa/dao"
	"git.gumpcome.com/go_kit/dbkit"
)

// @Title 添加账号
func AddAccountLogic(account *models.Account) (bool, int64, error) {

	// 1. 检查参数
	if account.RoleID == 0 || account.LoginName == "" || account.LoginPwd == "" || account.UserName == "" {
		// 参数不合法
		return false, 0, constant.RESP_CODE_PARAMS_ERROR
	}

	// 2. 准备数据
	params := map[string]interface{}{
		"user_name": account.UserName,
		"login_name": account.LoginName,
		"login_pwd": account.LoginPwd,
		"role_id": account.RoleID,
	}

	// 3. 添加权限
	return dao.AddAccountDao(&params)
}


// @Title 删除账号
func DeleteAccountLogic(account *models.Account) (bool, error) {

	// 1. 检查参数
	if account.ID == 0 {
		// 参数不合法
		return false, constant.RESP_CODE_PARAMS_ERROR
	}

	// 2. 准备数据
	params := map[string]interface{}{
		"id": account.ID,
	}

	// 3. 删除权限
	return dao.DeleteAccountDao(&params)
}


// @Title 分页查找账号
func PageAccountLogic(filter *models.PageAccountFilter) (dbkit.Page, error)  {
	return dao.PageAccountDao(filter)
}