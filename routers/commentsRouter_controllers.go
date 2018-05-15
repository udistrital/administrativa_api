package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActaRecibidoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEconomicaNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmortizacionElementoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AmparoContractualController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AsignacionPensionUserRoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:AtributoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:BodegaSalidaBodegaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CambioVidaUtilController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:CatalogoElementoGrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoElementoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DependenciaNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DetalleServicioNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadApropiacionSolicitudRpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:DocumentoSoporteActaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EspecificacionTecnicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoActaRecibidoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoInventarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:EstadoRolPensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:FuenteFinanciacionRubroNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:GrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:InventarioMovimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:IpcController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MarcoLegalNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModalidadSeleccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:MovimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NecesidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:OpcionAtributoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroContraloriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PeriodoPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PersonaEscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PolizaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetAllContratado",
			Router: `/Contratado/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion/:idPersona`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RangoEdadUpcController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoAtributoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionContratoEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroContraloriaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RelacionPolizaAmparoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ReposicionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:RequisitoMinimoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "ResolucionTemplate",
			Router: `/ResolucionTemplate/:dedicacion/:nivel`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "CancelarResolucion",
			Router: `/CancelarResolucion/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "RestaurarResolucion",
			Router: `/RestaurarResolucion/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GenerarResolucion",
			Router: `/GenerarResolucion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAllAprobada",
			Router: `/Aprobada`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAllExpedidasVigenciaPeriodo",
			Router: `/expedidas_vigencia_periodo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAllExpedidasVigenciaPeriodoVinculacion",
			Router: `/expedidas_vigencia_periodo_vinculacion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaBodegaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SalidaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "AddSolicitudRpTr",
			Router: `/AddSolicitudRpTr`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SolicitudRpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoCatalogoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubgrupoSubgrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SubtipoMovimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoBienController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoContratoNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoDocumentoSoporteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoEntradaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoFinanciacionNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNecesidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoUpcController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TipoValorAmparoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrAprobacionMasivaDocumentosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrAprobacionMasivaDocumentosController"],
		beego.ControllerComments{
			Method: "AprobarDocumentos",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrAprobacionMasivaPagosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrAprobacionMasivaPagosController"],
		beego.ControllerComments{
			Method: "AprobarPagos",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrAprobacionMasivaSoportesContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrAprobacionMasivaSoportesContratistasController"],
		beego.ControllerComments{
			Method: "AprobarSoportesContratistas",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNecesidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNecesidadDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:TrNecesidadDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UpcAdicionalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ValorizacionElementoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VerificacionNicspController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VigenciaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VigenciaContratoController"],
		beego.ControllerComments{
			Method: "VigenciaContrato",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetTotalContratosXResolucion",
			Router: `/get_total_contratos_x_resolucion/:id_resolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "InsertarVinculaciones",
			Router: `/InsertarVinculaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetVinculacionesAgrupadas",
			Router: `/get_vinculaciones_agrupadas/:id_resolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetValoresTotalesPorDisponibilidad",
			Router: `/get_valores_totales_x_disponibilidad/:anio/:periodo/:id_disponibilidad`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/administrativa_crud_api/controllers:ZonaUpcController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
