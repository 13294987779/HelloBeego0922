package db_mysql

import (
	"HelloBeegoProject/models"
	"NewMovieSpider/db_mysql"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
)
"github"

func Connect()  {
	config := beego.AppConfig//定义config变量，接收并赋值为全局变量
	appname := config.String("appname")
	fmt.Println("项目名为",appname)
	http,err:=config.Int("httpport")
	if err != nil{

		panic("项目配置信息解析错误，请查验后重试")
	}
	Db = db
	fmt.Println(db)
}
func Adduser(u models.User)(int64,error)  {
	md5Hash := md5.New()
	md5Hash.Write([]byte)


	result ,err :=Db.Exec
	if err !=nil {
		return -1,err
	}
	row,err :=result.RowsAffected()
	if err !=nil{
		result -1,err
	}




}