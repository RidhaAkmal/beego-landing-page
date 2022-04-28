package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.Data["Title"] = "Page About"
	c.TplName = "about/content.tpl"
}
