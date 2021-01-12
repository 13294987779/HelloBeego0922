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
func (c *MainController) Post(){

	//1、解析前端提交的json格式的数据
	var person models.Person
	dataBytes, err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("姓名:",person.Name)
	fmt.Println("年龄：",person.Age)
	fmt.Println("性别:",person.Sex)
	c.Ctx.WriteString("数据解析成功")
}

/**
 * 该方法用于处理delete请求
 */
func (c *MainController) Delete(){

}


