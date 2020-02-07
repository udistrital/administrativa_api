package controllers

import (
	"github.com/udistrital/administrativa_crud_api/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// operations for TrAprobacionMasivaDocumentosController
type TrAprobacionMasivaDocumentosController struct {
	beego.Controller
}

func (c *TrAprobacionMasivaDocumentosController) URLMapping() {
	c.Mapping("AprobarDocumentos", c.AprobarDocumentos)
}



// AprobarDocumentos ...
// @Title AprobarDocumentos
// @Description create AprobarDocumentos
// @Success 200 
// @Failure 403 body is empty
// @router / [post]
func (c *TrAprobacionMasivaDocumentosController) AprobarDocumentos() {

	var v []models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {


		if err = models.AprobarDocumentos(&v); err == nil {
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
