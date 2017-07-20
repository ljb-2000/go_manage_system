package models

//通用不带结果集的响应协议
type CommonResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//通用带结果集的响应协议
type CommonWithDataResp struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
