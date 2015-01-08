package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) Get() {
	/*QueryUnescape:使字符串为UTF-8格式
	RequestURI[1:]去除url前面的第一个/
	如http://localhost:8080/topic/view/6其url为/topic/view/6通过RequestURI[1:]变成topic/view/6
	以/开头容易以为是绝对*/
	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	if err != nil {
		//写文件
		this.Ctx.WriteString(err.Error())
		return
	}
	f, err := os.Open(filePath)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()
	/*https://gowalker.org/io#Copy
	io.Copy(dst, src)
	输出流dst writer
	输入流src reader */
	_, err = io.Copy(this.Ctx.ResponseWriter, f)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}
