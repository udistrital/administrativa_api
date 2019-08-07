package controllers

import (
  "github.com/udistrital/administrativa_crud_api/models"
  "github.com/astaxie/beego"
)

type PersonaEscalafonController struct {
	beego.Controller
}

func (c *PersonaEscalafonController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)

}

// GetAll ...
// @Title Get All
// @Description get PersonaEscalafon
// @Success 200 {object} models.PersonaEscalafon
// @Failure 403
// @router / [get]
func (c *PersonaEscalafonController) GetAll() {
    listaPersonas := models.GetAllPersonaEscalafon()
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaPersonas
    c.ServeJSON()
}
