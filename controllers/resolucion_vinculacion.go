package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_crud_api/models"
)

type ResolucionVinculacionController struct {
	beego.Controller
}

func (c *ResolucionVinculacionController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetAllAprobada", c.GetAllAprobada)
	c.Mapping("GetAllExpedidasVigenciaPeriodo", c.GetAllExpedidasVigenciaPeriodo)
	c.Mapping("GetAllExpedidasVigenciaPeriodoVinculacion", c.GetAllExpedidasVigenciaPeriodoVinculacion)
}

// GetAll ...
// @Title Get All
// @Description get ResolucionVinculacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ResolucionVinculacionDocente
// @Failure 403
// @router / [get]
func (c *ResolucionVinculacionController) GetAll() {
	listaResoluciones := models.GetAllResolucionVinculacion()
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = listaResoluciones
	c.ServeJSON()
}

// GetAllAprobada ...
// @Title Get All
// @Description get ResolucionVinculacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ResolucionVinculacionDocente
// @Failure 403
// @router /Aprobada [get]
func (c *ResolucionVinculacionController) GetAllAprobada() {
	listaResoluciones := models.GetAllResolucionAprobada()
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = listaResoluciones
	c.ServeJSON()
}

// GetAllExpedidasVigenciaPeriodoa ...
// @Title GetAllExpedidasVigenciaPeriodo
// @Description Agrupa los contratos de una preliquidacion segun mes, año y nomina para preliquidaicones en estado CERRADA
// @Param vigencia query string false "nomina a listar"
// @Param periodo query string false "mes de la liquidacion a listar"
// @Success 201 {object} models.Preliquidacion_x_contratos
// @Failure 403 body is empty
// @router /expedidas_vigencia_periodo [get]
func (c *ResolucionVinculacionController) GetAllExpedidasVigenciaPeriodo() {

	vigencia, err1 := c.GetInt("vigencia")
	periodo, err2 := c.GetInt("periodo")
	if err1 == nil && err2 == nil {

		listaResoluciones := models.GetAllExpedidasVigenciaPeriodo(vigencia, periodo)

		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = listaResoluciones

	} else {
		fmt.Println(err1)
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

// GetAllExpedidasVigenciaPeriodoa ...
// @Title GetAllExpedidasVigenciaPeriodo
// @Description Muestra resoluciones de tipo vinculación para cancelar y modificar
// @Param vigencia query string false "nomina a listar"
// @Param periodo query string false "mes de la liquidacion a listar"
// @Success 201 {object} models.Preliquidacion_x_contratos
// @Failure 403 body is empty
// @router /expedidas_vigencia_periodo_vinculacion [get]
func (c *ResolucionVinculacionController) GetAllExpedidasVigenciaPeriodoVinculacion() {

	vigencia, err1 := c.GetInt("vigencia")
	periodo, err2 := c.GetInt("periodo")
	if err1 == nil && err2 == nil {

		listaResoluciones := models.GetAllExpedidasVigenciaPeriodoVinculacion(vigencia, periodo)

		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = listaResoluciones

	} else {
		fmt.Println(err1)
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}
