package models

type Result struct {
	Code int `json:"code"`//接口返回状态类型 1表示成功  0表示失败
	Message string `json:"message"`//接口返回状态对应的描述信息
	Data interface{} `json:"data"`//返回的数据

}
