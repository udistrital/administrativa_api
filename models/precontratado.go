package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Precontratado struct {
	Id                 int     `orm:"column(id);pk;auto"`
	PrimerNombre       string  `orm:"column(primer_nombre)"`
	SegundoNombre      string  `orm:"column(segundo_nombre)"`
	PrimerApellido     string  `orm:"column(primer_apellido)"`
	SegundoApellido    string  `orm:"column(segundo_apellido)"`
	Documento          int     `orm:"column(documento)"`
	Expedicion         string  `orm:"column(expedicion)"`
	Categoria          string  `orm:"column(categoria)"`
	Dedicacion         string  `orm:"column(dedicacion)"`
	HorasSemanales     int     `orm:"column(horas_semanales)"`
	Semanas            int     `orm:"column(semanas)"`
	ProyectoCurricular int     `orm:"column(proyecto_curricular)"`
	ValorContrato      float64 `orm:"column(valor_contrato)"`
}

func init() {
	orm.RegisterModel(new(Precontratado))
}

func GetAllPrecontratado(idResolucion string) (arregloIDs []Precontratado) {
	o := orm.NewOrm()
	var temp []Precontratado
	_, err := o.Raw("SELECT vd.id id, ipn.primer_nombre, ipn.segundo_nombre, ipn.primer_apellido, ipn.segundo_apellido, ipn.num_documento_persona documento, c.nombre expedicion, e.nombre_escalafon categoria, d.nombre_dedicacion dedicacion, vd.numero_horas_semanales horas_semanales, vd.numero_semanas semanas, vd.id_proyecto_curricular FROM administrativa.vinculacion_docente vd INNER JOIN agora.informacion_persona_natural ipn ON ipn.num_documento_persona = vd.id_persona INNER JOIN agora.informacion_proveedor ip ON ipn.num_documento_persona = ip.num_documento INNER JOIN core.ciudad c ON ipn.id_ciudad_expedicion_documento = c.id_ciudad INNER JOIN administrativa.escalafon_persona ep ON ip.id_proveedor = ep.id_persona_natural INNER JOIN administrativa.escalafon e ON e.id_escalafon = ep.id_escalafon INNER JOIN administrativa.dedicacion d ON vd.id_dedicacion = d.id_dedicacion WHERE vd.estado = TRUE AND vd.id_resolucion = ? ;", idResolucion).QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}

func GetOnePrecontratado(idResolucion string, idPersona string) (precontratado *Precontratado) {
	o := orm.NewOrm()
	var temp []*Precontratado
	_, err := o.Raw("SELECT vd.id id, ipn.primer_nombre, ipn.segundo_nombre, ipn.primer_apellido, ipn.segundo_apellido, ipn.num_documento_persona documento, c.nombre expedicion, e.nombre_escalafon categoria, d.nombre_dedicacion dedicacion, vd.numero_horas_semanales horas_semanales, vd.numero_semanas semanas, vd.id_proyecto_curricular FROM administrativa.vinculacion_docente vd INNER JOIN agora.informacion_persona_natural ipn ON ipn.num_documento_persona = vd.id_persona INNER JOIN agora.informacion_proveedor ip ON ipn.num_documento_persona = ip.num_documento INNER JOIN core.ciudad c ON ipn.id_ciudad_expedicion_documento = c.id_ciudad INNER JOIN administrativa.escalafon_persona ep ON ip.id_proveedor = ep.id_persona_natural INNER JOIN administrativa.escalafon e ON e.id_escalafon = ep.id_escalafon INNER JOIN administrativa.dedicacion d ON vd.id_dedicacion = d.id_dedicacion WHERE vd.estado = TRUE AND vd.id_resolucion = ? AND ipn.num_documento_persona = ? ;", idResolucion, idPersona).QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	if temp != nil && len(temp) > 0 {
		return temp[0]
	} else {
		return nil
	}
}

func GetAllContratado(idResolucion string) (arregloIDs []Precontratado) {
	o := orm.NewOrm()
	var temp []Precontratado
	_, err := o.Raw("SELECT vd.id id, ipn.primer_nombre, ipn.segundo_nombre, ipn.primer_apellido, ipn.segundo_apellido, ipn.num_documento_persona documento, c.nombre expedicion, e.nombre_escalafon categoria, d.nombre_dedicacion dedicacion, vd.numero_horas_semanales horas_semanales, vd.numero_semanas semanas, vd.id_proyecto_curricular, cg.valor_contrato FROM administrativa.vinculacion_docente vd INNER JOIN agora.informacion_persona_natural ipn ON ipn.num_documento_persona = vd.id_persona INNER JOIN agora.informacion_proveedor ip ON ipn.num_documento_persona = ip.num_documento INNER JOIN core.ciudad c ON ipn.id_ciudad_expedicion_documento = c.id_ciudad INNER JOIN administrativa.escalafon_persona ep ON ip.id_proveedor = ep.id_persona_natural INNER JOIN administrativa.escalafon e ON e.id_escalafon = ep.id_escalafon INNER JOIN administrativa.dedicacion d ON vd.id_dedicacion = d.id_dedicacion INNER JOIN argo.contrato_general cg ON (vd.numero_contrato = cg.numero_contrato AND vd.vigencia = cg.vigencia) WHERE vd.estado = TRUE AND vd.id_resolucion = ? ;", idResolucion).QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}
