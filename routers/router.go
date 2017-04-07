// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/administrativa_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/necesidad_rechazada",
			beego.NSInclude(
				&controllers.NecesidadRechazadaController{},
			),
		),


			beego.NSNamespace("/solicitud_disponibilidad",
				beego.NSInclude(
					&controllers.SolicitudDisponibilidadController{},
				),
			),

			beego.NSNamespace("/necesidad_otro_si",
				beego.NSInclude(
					&controllers.NecesidadOtroSiController{},
				),
			),

			beego.NSNamespace("/actividad_especifica",
				beego.NSInclude(
					&controllers.ActividadEspecificaController{},
				),
			),

			beego.NSNamespace("/servicio",
				beego.NSInclude(
					&controllers.ServicioController{},
				),
			),

			beego.NSNamespace("/necesidad",
				beego.NSInclude(
					&controllers.NecesidadController{},
				),
			),

			beego.NSNamespace("/dependencia_necesidad",
				beego.NSInclude(
					&controllers.DependenciaNecesidadController{},
				),
			),

			beego.NSNamespace("/modalidad_seleccion",
				beego.NSInclude(
					&controllers.ModalidadSeleccionController{},
				),
			),

			beego.NSNamespace("/tipo_fuente_financiacion",
				beego.NSInclude(
					&controllers.TipoFuenteFinanciacionController{},
				),
			),

			beego.NSNamespace("/actividad_solicitud_necesidad",
				beego.NSInclude(
					&controllers.ActividadSolicitudNecesidadController{},
				),
			),

			beego.NSNamespace("/estado_necesidad",
				beego.NSInclude(
					&controllers.EstadoNecesidadController{},
				),
			),

			beego.NSNamespace("/requisito_minimo",
				beego.NSInclude(
					&controllers.RequisitoMinimoController{},
				),
			),

			beego.NSNamespace("/marco_legal",
				beego.NSInclude(
					&controllers.MarcoLegalController{},
				),
			),

			beego.NSNamespace("/marco_legal_necesidad",
				beego.NSInclude(
					&controllers.MarcoLegalNecesidadController{},
				),
			),

			beego.NSNamespace("/supervisor_solicitud_necesidad",
				beego.NSInclude(
					&controllers.SupervisorSolicitudNecesidadController{},
				),
			),

			beego.NSNamespace("/actividad_economica_necesidad",
				beego.NSInclude(
					&controllers.ActividadEconomicaNecesidadController{},
				),
			),

			beego.NSNamespace("/especificacion_tecnica",
				beego.NSInclude(
					&controllers.EspecificacionTecnicaController{},
				),
			),

			beego.NSNamespace("/fuente_financiacion_rubro_necesidad",
				beego.NSInclude(
					&controllers.FuenteFinanciacionRubroNecesidadController{},
				),
			),

			beego.NSNamespace("/tr_necesidad",
						beego.NSInclude(
							&controllers.TrNecesidadController{},
						),
					),
	)
	beego.AddNamespace(ns)
}
