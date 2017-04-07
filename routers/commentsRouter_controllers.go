package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			Method: "Acta_contrato",
			Router: `actaContrato/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/argo_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
