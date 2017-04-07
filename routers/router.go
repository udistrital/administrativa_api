package routers

import (
	"github.com/udistrital/administrativa_crud_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.ContratoGeneralController{},
			),
		),
		beego.NSNamespace("/acta_inicio",
			beego.NSInclude(
				&controllers.ActaInicioController{},
			),
		),
		beego.NSNamespace("/solicitud_rp",
			beego.NSInclude(
				&controllers.SolicitudRpController{},
			),
		),
		beego.NSNamespace("/disponibilidad_apropiacion_solicitud_rp",
			beego.NSInclude(
				&controllers.Disponibilidad_apropiacion_solicitud_rpController{},
			),
		),
		beego.NSNamespace("/acta_contrato",
			beego.NSInclude(
				&controllers.ActaInicioContratoGeneralController{},
			),
		),
		beego.NSNamespace("/ordenadores",
			beego.NSInclude(
				&controllers.ArgoOrdenadoresController{},
			),
		),
		beego.NSNamespace("/parametros",
			beego.NSInclude(
				&controllers.ParametrosController{},
			),
		),
		beego.NSNamespace("/parametro_estandar",
			beego.NSInclude(
				&controllers.ParametroEstandarController{},
			),
		),
		beego.NSNamespace("/relacion_parametro",
			beego.NSInclude(
				&controllers.RelacionParametroController{},
			),
		),
		beego.NSNamespace("/informacion_proveedor",
			beego.NSInclude(
				&controllers.InformacionProveedorController{},
			),
		),
		beego.NSNamespace("/supervisor",
			beego.NSInclude(
				&controllers.SupervisorContratoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
