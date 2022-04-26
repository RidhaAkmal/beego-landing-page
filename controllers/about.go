package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.TplName = "about.tpl"
}
