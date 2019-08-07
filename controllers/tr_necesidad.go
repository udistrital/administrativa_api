package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_crud_api/models"
)

// oprations for Necesidad
type TrNecesidadController struct {
	beego.Controller
}

func (c *TrNecesidadController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title TrNecesidad
// @Description insert the TrNecesidad in the tables Necesidad,FuenteFinanciacionRubroNecesidad, MarcoLegalNecesidad, DependenciaNecesidad, ActividadEspecifica, ActividadEconomicaNecesidad, DetalleServicioNecesidad, EspecificacionTecnica, RequisitoMinimo
// @Param	body		body 	models.TrNecesidad	true	"body for TrNecesidad content"
// @Success 201 {object} msg
// @Failure 403 :id is not int
// @router / [post]
func (c *TrNecesidadController) Post() {
	var v models.TrNecesidad
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if id, err := models.AddTrNecesidad(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			v.Necesidad.Id = int(id)
			c.Data["json"] = models.Alert{Type: models.AlertSucess, Body: v}
		} else {
			c.Data["json"] = models.Alert{Type: models.AlertError, Body: err}
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the TrNecesidad it calculates the consecutive number and update the need
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Necesidad	true		"body for TrNecesidad content"
// @Success 200 {object} models.TrNecesidad
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TrNecesidadController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.TrNecesidad{Necesidad: &models.Necesidad{Id: id}}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTrNecesidadByID(&v); err == nil {
			c.Data["json"] = "Ok"
		} else {
			c.Data["json"] = models.Alert{Type: models.AlertError, Body: err}
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
