package routers

import (
	"HelloBeegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//用户注册


    beego.Router("/index", &controllers.MainController{})
}
