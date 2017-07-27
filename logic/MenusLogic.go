package logic

import (
	"git.gumpcome.com/gumpoa/util"
	"git.gumpcome.com/gumpoa/models"
	"encoding/json"
	"git.gumpcome.com/gumpoa/dao"
	"github.com/astaxie/beego"
)


// @Title 根据账号对应的权限生成菜单
func CreateMenusLogic(account *models.Account) (*models.Menus, error){

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


