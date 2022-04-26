package routers

import (
	"beego-blog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
}
