package controllers

import (
	"github.com/astaxie/beego"

	"log"
	"party/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// // 退出登录
	// if this.Input().Get("exit") == "true" {
	// 	this.Redirect("/", 302)
	// 	return
	// }

	//this.TplNames = "login.html"
}

func (this *LoginController) Post() {

	//获取表单信息
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")

	log.Println(uname, pwd)

	// 验证用户名及密码
	ok := models.CheckAccount(uname, pwd)
	log.Println("ok:", ok)
	if ok != true {
		return
	}

	this.Redirect("/", 302)
	return

}
