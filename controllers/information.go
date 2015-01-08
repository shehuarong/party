package controllers

import (
	"github.com/astaxie/beego"
	"party/models"
	"path"
)

type InformationController struct {
	beego.Controller
}

func (this *InformationController) Get() {
	// TODO:先检查是是否登录（未完成）
	this.Data["IsInformation"] = true
	this.TplNames = "information.html"
	users, err := models.GetAllUsers("", "")
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Users"] = users
	}

}
func (this *InformationController) Post() {
	// TODO:先检查是是否登录（未完成）

	// 解析表单
	uid := this.Input().Get("uid")

	username := this.Input().Get("username")
	sex := this.Input().Get("sex")
	identity := this.Input().Get("identity")
	native := this.Input().Get("native")
	nation := this.Input().Get("nation")
	number := this.Input().Get("number")
	education := this.Input().Get("education")
	phone := this.Input().Get("phone")
	email := this.Input().Get("email")
	address := this.Input().Get("address")
	resident := this.Input().Get("resident")
	// 从表单获取的时间为string类型
	birthday := this.Input().Get("birthday")
	jointime := this.Input().Get("jointime")
	company := this.Input().Get("company")
	position := this.Input().Get("position")
	group := this.Input().Get("group")
	branch := this.Input().Get("branch")
	category := this.Input().Get("category")
	linkman := this.Input().Get("linkman")
	introducer := this.Input().Get("introducer")
	feature := this.Input().Get("feature")
	comment := this.Input().Get("comment")

	// 获取照片
	_, fh, err := this.GetFile("photo")
	if err != nil {
		beego.Error(err)

	}
	var photo string
	if fh != nil {
		// 保存照片
		photo = fh.Filename
		beego.Info(photo)
		/*path.Join()将两个字符串拼接起来，根据系统的不同，有不同的分隔符
		filename:tmp.go
		attachment/tmp.go
		*/
		err := this.SaveToFile("photo", path.Join("photo", photo))
		if err != nil {
			beego.Error(err)
		}
		if len(username) == 0 {
			err = models.AddInformation(uid, username, sex, identity, native, nation, number, education, phone, email, address, resident, company, position, group, branch, category, linkman, introducer, feature, comment, photo, birthday, jointime)

		} else {
			err = models.ModifyUser(uid, username, sex, identity, native, nation, number, education, phone, email, address, resident, company, position, group, branch, category, linkman, introducer, feature, comment, photo, birthday, jointime)
		}
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/information", 302)

}
func (this *InformationController) Add() {
	this.TplNames = "information_add.html"
	uid := this.Input().Get("uid")
	user, err := models.GetUser(uid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	// 获取数据
	this.Data["User"] = user
	this.Data["Uid"] = uid
}

func (this *InformationController) View() {
	this.TplNames = "information_view.html"
	user, err := models.GetUser(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["User"] = user
	this.Data["Uid"] = this.Ctx.Input.Param("0")
}
func (this *InformationController) Modify() {
	this.TplNames = "information_modify.html"
	uid := this.Input().Get("uid")
	user, err := models.GetUser(uid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	// 获取数据
	this.Data["User"] = user
	this.Data["Uid"] = uid

}
func (this *InformationController) Delete() {
	err := models.DeleteUser(this.Input().Get("uid"))
	if err != nil {
		beego.Error(err)

	}
	this.Redirect("/", 302)
	return
}
