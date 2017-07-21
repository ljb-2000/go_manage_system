package models

//项目内应用的所有业务实体

//登录信息
type UserInfo struct {
	Account  string `json:"account" form:"account" desc:"账号"`
	Password string `json:"password" form:"password" desc:"密码"`
}

//权限资源
type Access struct {
	ID         int 		`json:"id" form:"id" desc:"序号"`
	Code       int 		`json:"code" form:"code" desc:"权限码"`
	Name       string 	`json:"name" form:"name" desc:"权限名"`
	CreateTime string 	`json:"create_time" form:"create_time" desc:"创建时间"`
}
