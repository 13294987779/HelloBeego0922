package main

import (
	_ "HelloBeegoProject/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	config := beego.AppConfig//定义config变量，接收并赋值为全局变量
	appname := config.String("appname")
	fmt.Println("项目名为",appname)
	http,err:=config.Int("httpport")
	if err != nil{
		fmt.Println("")
		panic("项目配置信息解析错误，请查验后重试")
	}
	fmt.Println("http端口号为",http)
	driver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPwd := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	db,err := sql.Open(driver,dbUser+":"+dbPwd+"@tpc("+dbIp+")/"+dbName+"?charset=utf8")
	if err != nil{
		fmt.Println("错误")
	}
	fmt.Println(db)
	beego.Run()


}


