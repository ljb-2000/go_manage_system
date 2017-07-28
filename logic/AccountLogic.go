package logic

import (
	"encoding/json"
	"git.fenggese.com/go_kit/dbkit"
	"git.fenggese.com/go_manage_system/constant"
	"git.fenggese.com/go_manage_system/dao"
	"git.fenggese.com/go_manage_system/models"
	"git.fenggese.com/go_manage_system/util"
	"github.com/astaxie/beego"
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
		"user_name":  account.UserName,
		"login_name": account.LoginName,
		"login_pwd":  account.LoginPwd,
		"role_id":    account.RoleID,
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
func PageAccountLogic(filter *models.PageAccountFilter) (dbkit.Page, error) {
	// 1. 检查参数
	if filter.PageNumber < 1 || filter.PageSize < 0 {
		// 参数不合法
		return nil, constant.RESP_CODE_PARAMS_VALUE_ERROR
	}

	// 2. 分页查找
	return dao.PageAccountDao(filter)
}

// @Title 验证用户登录信息
func LoginLogic(user *models.UserInfo) (bool, error) {

	// 1. 准备数据
	params := map[string]interface{}{
		"login_name": user.Account,
		"login_pwd":  user.Password,
	}

	// 2. 查询数据库
	count, err := dao.LoginDao(&params)
	if err != nil {
		// 数据库查询出错
		return false, err
	}

	// 3. 处理结果
	if count != 1 {
		// 登录失败
		return false, nil
	}
	// 登录成功
	return true, nil

}

// @Title 查询账号的权限
func AccessQueryByAccountLogic(account *models.Account) ([]map[string]interface{}, error) {
	// 1. 检查参数
	if account.LoginName == "" {
		return nil, constant.RESP_CODE_PARAMS_VALUE_ERROR
	}

	// 2. 开始查找
	return dao.AccessQueryByAccountDao(account.LoginName)
}

// @Title 根据账号对应的权限生成菜单
func CreateMenusLogic(account *models.Account) (*models.Menus, error) {

	// 1. 读取JSON文件 相对地址
	bytes, err := util.ReadFile("./models/menus.json")
	if err != nil {
		return nil, err
	}

	// 2. 解析JSON
	var menus models.Menus
	err = json.Unmarshal(bytes, &menus)
	if err != nil {
		return nil, err
	}

	// 如果是管理员，获得所有菜单
	if account.LoginName == beego.AppConfig.String("admin_account") {
		return &menus, nil
	}

	// 3. 获取账号的权限列表
	accesses, err := dao.AccessQueryByAccountDao(account.LoginName)
	if err != nil {
		return nil, err
	}

	// 4. 根据权限生成菜单
	var newMenus models.Menus
	for _, access := range accesses {
		for _, sub := range menus.Subs {
			if sub.Code == access["code"] {
				newMenus.Subs = append(newMenus.Subs, sub)
			}
		}
	}

	return &newMenus, nil

}

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
