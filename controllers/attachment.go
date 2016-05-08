package controllers

import (
	//"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type AttachController struct {
	BaseController
}

func (this *AttachController) Get() {
	//uri_orign := this.Ctx.Request.RequestURI
	uri := this.Ctx.Request.RequestURI[1:]
	//filePath := url.QueryEscape(uri)
	filePath, err := url.QueryUnescape(uri)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	// beego.Info("orign:" + uri_orign)
	// beego.Info("second:" + uri)
	// beego.Info("final:" + filePath)
	f, err := os.Open(filePath)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(this.Ctx.ResponseWriter, f)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}
