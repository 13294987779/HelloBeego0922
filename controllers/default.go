package controllers

import (
	"HelloBeegoProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	name := c.GetString("name")
	//age,err:= c.GetInt("age",)
	//获取请求参数
	//name := c.Ctx.Input.Query("name")
	age :=  c.Ctx.Input.Query("age")
	sex :=  c.Ctx.Input.Query("sex")
	////fmt.Println(name,age,sex)

	if name != "dan" || age != "18" || sex != "man"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("数据校验错误，登录失败"))
		fmt.Println("用户登录失败")
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据校验正确，登录成功"))
	fmt.Println("用户名为",name)
	fmt.Println("年龄为",age)
	fmt.Println("性别为",sex)
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "971371682@qq.com"
	c.TplName = "index.tpl"

	}

/**
*该post方法是处理post
 */
func (c *MainController)Post()  {
	//该post方法会处理post类型请求时调用的方法
	//fmt.Println("post类型的请求")
	//username :=  c.Ctx.Request.FormValue("username")
	//fmt.Println("用户名为",username)
	//pwd :=  c.Ctx.Request.FormValue("pwd")
	//fmt.Println("密码为",pwd)
	body := c.Ctx.Request.Body
	dataBytes,err :=  ioutil.ReadAll(body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败")
		return
	}
	//json解析
	var persion models.Man
	json.Unmarshal(dataBytes,persion)
	//fmt.Println("用户名",persion.Name,"年龄",persion.Age,"性别",persion.Sex)


	//条件判断
	//if username != "admin" || pwd != "123456"{
	//	//失败页面
	//	c.Ctx.ResponseWriter.Write([]byte("对不起，登录失败"))
	//	return
	//}
	//c.Ctx.ResponseWriter.Write([]byte("登录成功"))
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "971371682@qq.com"
	c.TplName = "index.tpl"

}


