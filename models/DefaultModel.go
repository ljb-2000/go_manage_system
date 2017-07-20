package models


const BusinessWrong int = 10003


type Result map[string]interface{}


func NewResult() Result{
	return Result{
		"data": map[string]interface{}{},
		"msg": "响应成功",
		"code": 200,
	}
}
