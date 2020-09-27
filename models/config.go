package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"golang.org/x/text/language"
)
func main(){
	config :=beego.AppConfig
	appName :=config.String("appname")
	fmt.Println("项目应用名称：",appName)
	if err != nil{
		panic("")
	}
	fmt.Println("",port)
	beego.Run()
}
