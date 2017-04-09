package controllers

import (
	"encoding/json"
	"github.com/udistrital/administrativa_crud_api/models"
	"github.com/astaxie/beego"
)

// oprations for Necesidad
type TrNecesidadController struct {
	beego.Controller
}

func (c *TrNecesidadController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title TrNecesidad
// @Description insert the TrNecesidad
// @Param	body		body 	models.TrNecesidad	true	"body for TrNecesidad content"
// @Success 200 {object} msg
// @Failure 403 :id is not int
// @router / [post]
func (c *TrNecesidadController) Post() {
	var v models.TrNecesidad
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alerta, err := models.AddTrNecesidad(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alerta
		} else {
			c.Data["json"] = alerta
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()

}
