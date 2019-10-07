package controllers

import (
  "github.com/udistrital/administrativa_crud_api/models"
  "github.com/astaxie/beego"
)

type PrecontratadoController struct {
	beego.Controller
}

func (c *PrecontratadoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
    c.Mapping("GetAllContratado", c.GetAllContratado)
	c.Mapping("GetOne", c.GetOne)
}

// GetAll ...
// @Title Get All
// @Description get Precontratado
// @Param   idResolucion      path    string  true        "The key for staticblock"
// @Success 200 {object} models.Precontratado
// @Failure 403
// @router /:idResolucion [get]
func (c *PrecontratadoController) GetAll() {
	idResolucion := c.Ctx.Input.Param(":idResolucion")
    listaPrecontratados := models.GetAllPrecontratado(idResolucion)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaPrecontratados
    c.ServeJSON()
}

// GetAllContratado ...
// @Title Get All
// @Description get Precontratado
// @Param   idResolucion      path    string  true        "The key for staticblock"
// @Success 200 {object} models.Precontratado
// @Failure 403
// @router /Contratado/:idResolucion [get]
func (c *PrecontratadoController) GetAllContratado() {
    idResolucion := c.Ctx.Input.Param(":idResolucion")
    listaPrecontratados := models.GetAllContratado(idResolucion)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaPrecontratados
    c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Precontratado by id
// @Param   idResolucion     path    string  true        "The key for staticblock"
// @Param   idPersona        path    string  true        "The key for staticblock"
// @Success 200 {object} models.Precontratado
// @Failure 403 :id is empty
// @router /:idResolucion/:idPersona [get]
func (c *PrecontratadoController) GetOne() {
	idResolucion := c.Ctx.Input.Param(":idResolucion")
	idPersona := c.Ctx.Input.Param(":idPersona")
    precontratado := models.GetOnePrecontratado(idResolucion, idPersona)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = precontratado
    c.ServeJSON()
}
