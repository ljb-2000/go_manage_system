package logic

// 左侧菜单
type SubMenu struct {
	Code     int 		`json:"code" desc:"权限码"`
	Name     string     `json:"name" desc:"菜单名"`
	Path     string     `json:"path" desc:"路径名"`
	IconCls  string     `json:"iconCls" desc:"css的样式"`
	Children []MenuItem `json:"children" desc:"菜单项"`
}

// 菜单项
type MenuItem struct {
	Name string `json:"name" desc:"权限名／菜单名"`
	Path string `json:"path" desc:"路径"`
}

// @Title 生成菜单
// @Description 还没有查询数据库，账号管理模块完成后还要修改
func CreateMenusLogic(account string) (*[]SubMenu, bool) {

	// 账户管理 菜单项
	account_menu_items := []MenuItem{
		{"权限管理", "/access"},
		{"角色管理", "/role"},
		{"账号管理", "/account"},
	}
	// 账户管理
	accout_menu := SubMenu{
		101,
		"账户管理",
		"/",
		"el-icon-message",
		account_menu_items,
	}

	// 财务结算菜单项（占位用的）
	settlement_menu_items := []MenuItem{
		{"财务管理", "#"},
		{"结算管理", "#"},
	}

	// 财务结算
	settlement_menu := SubMenu{
		102,
		"财务管理",
		"/",
		"el-icon-document",
		settlement_menu_items,
	}

	// 左侧菜单
	menus := []SubMenu{
		accout_menu,
		settlement_menu,
	}

	return &menus, true
}
