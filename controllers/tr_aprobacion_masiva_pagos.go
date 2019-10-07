package controllers

import (
	"github.com/udistrital/administrativa_crud_api/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// operations for TrAprobacionMasivaPagosController
type TrAprobacionMasivaPagosController struct {
	beego.Controller
}

func (c *TrAprobacionMasivaPagosController) URLMapping() {
	c.Mapping("AprobarPagos", c.AprobarPagos)
}



// AprobarPagos ...
// @Title AprobarPagos
// @Description create AprobarPagos
// @Success 200 
// @Failure 403 body is empty
// @router / [post]
func (c *TrAprobacionMasivaPagosController) AprobarPagos() {

	var v []models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {


		if err = models.AprobarPagos(&v); err == nil {
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
