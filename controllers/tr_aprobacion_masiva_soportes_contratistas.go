package controllers

import (
	"github.com/udistrital/administrativa_crud_api/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// operations for TrAprobacionMasivaSoportesContratistasController
type TrAprobacionMasivaSoportesContratistasController struct {
	beego.Controller
}

func (c *TrAprobacionMasivaSoportesContratistasController) URLMapping() {
	c.Mapping("AprobarSoportesContratistas", c.AprobarSoportesContratistas)
}



// AprobarSoportesContratistas ...
// @Title AprobarSoportesContratistas
// @Description create AprobarSoportesContratistas
// @Success 200 
// @Failure 403 body is empty
// @router / [post]
func (c *TrAprobacionMasivaSoportesContratistasController) AprobarSoportesContratistas() {

	var v []models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {


		if err = models.AprobarSoportesContratistas(&v); err == nil {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = v

		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
