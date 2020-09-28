package controllers

import (
	"HelloBeegoProject/db_mysql"
	"HelloBeegoProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type RegisterController struct {
	beego.Controller
}

func (r*RegisterController)post()  {
	dataBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	if err !=nil {
		r.Ctx.WriteString("数据接收错误")
		return
	}
	var user models.User
	err =json.Unmarshal(dataBytes,&user)
	if err !=nil{
		r.Ctx.WriteString("数据解析错误，请重试")
		return
	}
row, err :=db_mysql.Adduser(user)
	if err !=nil {
		 r.Ctx.WriteString("注册用户信息失败")
		return
	}
	fmt.Println(row)
r.Ctx.WriteString("注册成功")
http.HandleFunc("/")
}
