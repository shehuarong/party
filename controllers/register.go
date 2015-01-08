package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	"party/models"
	"strings"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.Data["IsRegister"] = true
	this.TplNames = "register.html"
	nickname := this.Input().Get("nickname")
	// email := this.Input().Get("email")
	password := this.Input().Get("password")
	password2 := this.Input().Get("password2")
	ok := strings.EqualFold(password, password2)
	// if !ok {
	// 	return
	// }
	if len(nickname) == 0 && len(password) == 0 && len(password2) == 0 && ok {
		return
	}

	err := models.AddUser(nickname, password)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/login.html", 302)
	return

}
