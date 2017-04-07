package routers

import (
	"github.com/astaxie/beego"
)

func init() {

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
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:ServicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

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

}
