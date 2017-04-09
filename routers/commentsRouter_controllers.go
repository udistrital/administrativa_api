package routers

import (
	"github.com/astaxie/beego"
)

func init() {

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			Method: "Acta_contrato",
			Router: `actaContrato/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ActaInicioController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ArgoOrdenadoresController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ContratoGeneralController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:InformacionProveedorController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametroEstandarController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ParametrosController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RelacionParametroController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudRpController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
=======
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorContratoController"],
>>>>>>> solicitud_rp
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TrNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:TrNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

=======
>>>>>>> solicitud_rp
}
