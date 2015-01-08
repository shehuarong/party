package controllers

import (
	"github.com/astaxie/beego"
)

type HonorController struct {
	beego.Controller
}

func (this *HonorController) Get() {
	this.Data["IsHonor"] = true
	this.TplNames = "honor.html"

}
