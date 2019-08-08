package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_crud_api/models"
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
// @Failure 404 not found resource
// @router / [get]
func (c *PersonaEscalafonController) GetAll() {
	listaPersonas := models.GetAllPersonaEscalafon()
	c.Ctx.Output.SetStatus(201)
	if listaPersonas == nil {
		listaPersonas = append(listaPersonas)
	}
	c.Data["json"] = listaPersonas
	c.ServeJSON()
}
