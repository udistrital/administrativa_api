package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func GetVigenciaContrato() (vigencias []int) {
	o := orm.NewOrm()
	var temp []int
	_, err := o.Raw("SELECT DISTINCT vigencia FROM argo.contrato_general ORDER BY vigencia DESC;").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	fmt.Println(temp)
	return temp
}
