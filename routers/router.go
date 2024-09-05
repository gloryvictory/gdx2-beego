// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/gloryvictory/gdx2-beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/lu",
			beego.NSInclude(
				&controllers.LuController{},
			),
		),

		beego.NSNamespace("/field",
			beego.NSInclude(
				&controllers.FieldController{},
			),
		),

		beego.NSNamespace("/stp",
			beego.NSInclude(
				&controllers.StpController{},
			),
		),

		beego.NSNamespace("/stl",
			beego.NSInclude(
				&controllers.StlController{},
			),
		),

		beego.NSNamespace("/sta",
			beego.NSInclude(
				&controllers.StaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
