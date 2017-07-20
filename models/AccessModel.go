package models

type Access struct {
	ID         int 		`json:"id" desc:"序号"`
	Name       string 	`json:"access_name" desc:"权限名"`
	Code       int 		`json:"access_code" desc:"权限码"`
	CreateTime string 	`json:"create_time" desc:"创建时间"`
}
