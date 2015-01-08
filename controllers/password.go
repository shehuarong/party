package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"party/models"
)

type PasswordController struct {
	beego.Controller
}

func (this *PasswordController) Get() {
	this.TplNames = "password_modify.html"
}

func (this *PasswordController) Post() {
	uid := this.Input().Get("uid")
	user, err := models.GetUser(uid)
	log.Println(user, err)
}
