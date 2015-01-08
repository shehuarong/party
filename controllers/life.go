package controllers

import (
	"github.com/astaxie/beego"
)

type LifeController struct {
	beego.Controller
}

func (this *LifeController) Get() {
	this.Data["IsLife"] = true
	this.TplNames = "life.html"

}
