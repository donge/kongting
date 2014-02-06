package routers

import (
	"github.com/astaxie/beego"
	"kongting/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/weixin", &controllers.WeixinController{})
}
