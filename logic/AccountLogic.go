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

	// 3. 添加账号
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

	// 3. 删除账号
	return dao.DeleteAccountDao(&params)
}


// @Title 分页查找账号
func PageAccountLogic(filter *models.PageAccountFilter) (dbkit.Page, error)  {
	return dao.PageAccountDao(filter)
}

// @Title 验证用户登录信息
func LoginLogic(user *models.UserInfo) (bool, error) {
	account := models.Account{}
	account.LoginName = user.Account
	account.LoginPwd = user.Password

	// 查询数据库
	count, err := dao.LoginDao(&account)
	if err != nil {
		// 数据库查询出错
		return false, err
	}

	// 登录失败
	if count != 1 {
		return false, nil
	}
	// 登录成功
	return true, nil

}


// @Title 查询账号的权限
func AccessQueryByAccountLogic(account *models.Account) ([]map[string]interface{}, error) {
	if account.LoginName == "" {
		return nil, constant.RESP_CODE_ACCOUNT_ACCESSES_ERROR
	}
	return dao.AccessQueryByAccountDao(account.LoginName)
}
