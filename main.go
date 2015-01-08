package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"party/controllers"
	"party/models"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/password_modify", &controllers.PasswordController{})

	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/life", &controllers.LifeController{})

	beego.Router("/honor", &controllers.HonorController{})
	beego.Router("/information", &controllers.InformationController{})
	/*beego的自动路由
	http://beego.me/docs/mvc/controller/router.md
	找到自动匹配*/
	beego.AutoRouter(&controllers.InformationController{})
	beego.Router("/category", &controllers.CategoryController{})

	os.Mkdir("photo", os.ModePerm)
	beego.Router("/photo/:all", &controllers.PhotoController{})
	beego.Run()
}
