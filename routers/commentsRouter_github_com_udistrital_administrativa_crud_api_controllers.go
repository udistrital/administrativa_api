package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			"Acta_contrato",
			`actaContrato/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioContratoGeneralController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadSolicitudNecesidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:Disponibilidad_apropiacion_solicitud_rpController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadOtroSiController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadRechazadaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ServicioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorSolicitudNecesidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFuenteFinanciacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNecesidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VigenciaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VigenciaContratoController"],
		beego.ControllerComments{
			"VigenciaContrato",
			`/`,
			[]string{"get"},
			nil})

}
