package main

import (
	"HelloBeegoProject/db_mysql"
	_ "HelloBeegoProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//1、连接数据库
	db_mysql.Connect()

	//2、其他配置

	//3、启动应用
	beego.Run()//代码简洁
}
	/*config := beego.AppConfig//定义config变量，接收并赋值为全局变量
	appname := config.String("appname")
	fmt.Println("项目名为",appname)
	http,err:=config.Int("httpport")
	if err != nil{

		panic("项目配置信息解析错误，请查验后重试")
	}
	fmt.Println("http端口号为",http)
	dbDriver := config.String("db_drivername")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriver,dbUser,dbPassword)



	db,err :=dbUser+":"+dbPassword+"@tpc("+dbIp+")/"+dbName+"?charset=utf8"

	if err != nil{
		panic("数据库连接错误，请检查配置")
	}
	fmt.Println(db)
	beego.Run()

*/



