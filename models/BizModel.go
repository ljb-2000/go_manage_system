package models

//项目内应用的所有业务实体

//登录信息
type UserInfo struct {
	Account  string `json:"account" form:"account" desc:"账号"`
	Password string `json:"password" form:"password" desc:"密码"`
}

// *********************菜单***********************
// 菜单
type Menus struct {
	Subs []SubMenu `json:"menus" desc:"菜单列表"`
}

// 左侧菜单
type SubMenu struct {
	Code     int        `json:"code" desc:"权限码"`
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

// *********************权限***********************
// 权限
type Access struct {
	ID         int    `json:"id" form:"id" desc:"序号"`
	Code       int    `json:"code" form:"code" desc:"权限码"`
	Name       string `json:"name" form:"name" desc:"权限名"`
	CreateTime string `json:"create_time" form:"create_time" desc:"创建时间"`
}

// 权限分页查找
type PageAccessFilter struct {
	PageNumber int `form:"page_number" desc:"页码"`
	PageSize   int `form:"page_size" desc:"每页最大记录数"`
}

// *********************角色***********************
// 角色
type Role struct {
	ID         int    `json:"id" form:"id" desc:"序号"`
	Name       string `json:"name" form:"name" desc:"角色名"`
	CreateTime string `json:"create_time" form:"create_time" desc:"创建时间"`
}

// 角色权限
type RoleAccess struct {
	ID         int    `json:"id" form:"id" desc:"序号"`
	RoleID     int    `json:"role_id" form:"role_id" desc:"角色ID"`
	AccessID   int    `json:"access_id" form:"access_id" desc:"权限ID"`
	CreateTime string `json:"create_time" form:"create_time" desc:"创建时间"`
}

// 角色分页查找
type PageRoleFilter struct {
	PageNumber int `form:"page_number" desc:"页码"`
	PageSize   int `form:"page_size" desc:"每页最大记录数"`
}

// *********************账号***********************
type Account struct {
	ID         int    `json:"id" form:"id" desc:"序号"`
	LoginName  string `json:"login_name" form:"login_name" desc:"账号"`
	LoginPwd   string `json:"login_pwd" form:"login_pwd" desc:"密码"`
	UserName   string `json:"user_name" form:"user_name" desc:"用户名"`
	RoleID     int    `json:"role_id" form:"role_id" desc:"角色ID"`
	CreateTime string `json:"create_time" form:"create_time" desc:"创建时间"`
}

// 账号分页查找
type PageAccountFilter struct {
	PageNumber int `form:"page_number" desc:"页码"`
	PageSize   int `form:"page_size" desc:"每页最大记录数"`
}
