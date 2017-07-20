package models

//项目内应用的所有业务实体

//登录信息
type UserInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

//权限资源
type Access struct {
	ID         int 		`json:"id" desc:"序号"`
	Name       string 	`json:"access_name" desc:"权限名"`
	Code       int 		`json:"access_code" desc:"权限码"`
	CreateTime string 	`json:"create_time" desc:"创建时间"`
}
