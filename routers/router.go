// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"pruebaOAS/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/responsable",
			beego.NSInclude(
				&controllers.ResponsableController{},
			),
		),

		beego.NSNamespace("/estado",
			beego.NSInclude(
				&controllers.EstadoController{},
			),
		),

		beego.NSNamespace("/actividad",
			beego.NSInclude(
				&controllers.ActividadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
