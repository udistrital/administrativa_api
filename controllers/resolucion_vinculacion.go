package controllers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"

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
// @Failure 404 not found resource
// @router / [get]
func (c *ResolucionVinculacionController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	l, err := models.GetAllResolucionVinculacion(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l)
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// GetAllAprobada ...
// @Title Get All
// @Description get ResolucionVinculacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	int		false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	int		false	"Start position of result set. Must be an integer"
// @Success 201 {object} models.ResolucionVinculacionDocente
// @Failure 404 not found resource
// @router /Aprobada [get]
func (c *ResolucionVinculacionController) GetAllAprobada() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	listaResoluciones, err := models.GetAllResolucionAprobada(query, fields, sortby, order, offset, limit)
	if err != nil {
		beego.Error(err)
		c.Abort("403")
	}

	c.Ctx.Output.SetStatus(201)
	if listaResoluciones == nil {
		listaResoluciones = append(listaResoluciones)
	}
	c.Data["json"] = listaResoluciones
	c.ServeJSON()
}

// GetAllExpedidasVigenciaPeriodo ...
// @Title GetAllExpedidasVigenciaPeriodo
// @Description Agrupa los contratos de una preliquidacion segun mes, año y nomina para preliquidaicones en estado CERRADA
// @Param vigencia query string false "nomina a listar"
// @Success 201 {object} models.Preliquidacion_x_contratos
// @Failure 403 body is empty
// @router /expedidas_vigencia_periodo [get]
func (c *ResolucionVinculacionController) GetAllExpedidasVigenciaPeriodo() {

	vigencia, err := c.GetInt("vigencia")
	if err == nil {

		listaResoluciones := models.GetAllExpedidasVigenciaPeriodo(vigencia)

		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = listaResoluciones

	} else {
		fmt.Println(err)
		c.Abort("403")
	}
	c.ServeJSON()
}

// GetAllExpedidasVigenciaPeriodoVinculacion ...
// @Title GetAllExpedidasVigenciaPeriodoVinculacion
// @Description Muestra resoluciones de tipo vinculación para cancelar y modificar
// @Param vigencia query string false "nomina a listar"
// @Success 201 {object} models.Preliquidacion_x_contratos
// @Failure 403 body is empty
// @router /expedidas_vigencia_periodo_vinculacion [get]
func (c *ResolucionVinculacionController) GetAllExpedidasVigenciaPeriodoVinculacion() {

	vigencia, err := c.GetInt("vigencia")
	if err == nil {

		listaResoluciones := models.GetAllExpedidasVigenciaPeriodoVinculacion(vigencia)

		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = listaResoluciones

	} else {
		fmt.Println(err)
		c.Abort("403")
	}
	c.ServeJSON()
}
