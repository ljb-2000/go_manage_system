package models

// 登录信息
type UserInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}