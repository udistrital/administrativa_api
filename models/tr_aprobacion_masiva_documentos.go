package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//funcion para la aprobación masiva de documentos
func AprobarDocumentos(m *[]PagoMensual) (err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range *m {
		v.EstadoPagoMensual.Id = 3
		if _, err = o.Update(&v); err != nil {
			fmt.Println("Pago mensual documentos aprobados", &v)
			err = o.Rollback()
		} else {
			fmt.Println(err)
		}
	}
	err = o.Commit()

	return
}
