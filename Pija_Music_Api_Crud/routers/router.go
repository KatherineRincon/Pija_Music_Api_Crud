// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Trajes_Tipicos",
			beego.NSInclude(
				&controllers.TrajesTipicosController{},
			),
		),

		beego.NSNamespace("/Comida",
			beego.NSInclude(
				&controllers.ComidaController{},
			),
		),

		beego.NSNamespace("/Lugares",
			beego.NSInclude(
				&controllers.LugaresController{},
			),
		),

		beego.NSNamespace("/Cultura",
			beego.NSInclude(
				&controllers.CulturaController{},
			),
		),

		beego.NSNamespace("/Autor_Coplas",
			beego.NSInclude(
				&controllers.AutorCoplasController{},
			),
		),

		beego.NSNamespace("/Coplas",
			beego.NSInclude(
				&controllers.CoplasController{},
			),
		),

		beego.NSNamespace("/Credenciales",
			beego.NSInclude(
				&controllers.CredencialesController{},
			),
		),

		beego.NSNamespace("/Usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),

		beego.NSNamespace("/Favoritos",
			beego.NSInclude(
				&controllers.FavoritosController{},
			),
		),

		beego.NSNamespace("/Canciones",
			beego.NSInclude(
				&controllers.CancionesController{},
			),
		),

		beego.NSNamespace("/Estilo_Musical",
			beego.NSInclude(
				&controllers.EstiloMusicalController{},
			),
		),

		beego.NSNamespace("/Artista",
			beego.NSInclude(
				&controllers.ArtistaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
