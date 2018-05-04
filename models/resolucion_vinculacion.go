package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type ResolucionVinculacion struct {
	Id              int       `orm:"column(id);pk;auto"`
	Estado          string    `orm:"column(estado)"`
	Numero          string    `orm:"column(numero)"`
	Vigencia        int       `orm:"column(vigencia)"`
	Facultad        int       `orm:"column(facultad)"`
	NivelAcademico  string    `orm:"column(nivel_academico)"`
	Dedicacion      string    `orm:"column(dedicacion)"`
	FechaExpedicion time.Time `orm:"column(fecha_expedicion);type(timestamp without time zone)"`
	NumeroSemanas   int       `orm:"column(numero_semanas)"`
	Periodo         int       `orm:"column(periodo)"`
	TipoResolucion  string    `orm:"column(tipo_resolucion)"`
}

func init() {
	orm.RegisterModel(new(ResolucionVinculacion))
}

func GetAllResolucionVinculacion() (arregloIDs []ResolucionVinculacion) {
	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	_, err := o.Raw("SELECT DISTINCT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.id_facultad facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion, tipo.nombre_tipo_resolucion tipo_resolucion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, administrativa.resolucion_estado re, administrativa.estado_resolucion e, administrativa.tipo_resolucion tipo WHERE r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND r.id_tipo_resolucion=tipo.id_tipo_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND re.estado!=6 ORDER BY id desc;").QueryRows(&temp)

	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	for x, resoluciones := range temp {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		temp[x].FechaExpedicion = resoluciones.FechaExpedicion
	}
	return temp
}

func GetAllResolucionAprobada(limit, offset int, query string) (arregloIDs []ResolucionVinculacion, err error) {
	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	//_, err := o.Raw("SELECT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, d.nombre facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, oikos.dependencia d, administrativa.resolucion_estado re, administrativa.estado_resolucion e WHERE rv.id_facultad=d.id AND r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND r.id_tipo_resolucion=1 ORDER BY id desc;").QueryRows(&temp)
	//TODO: dar soporte a query (sin dejar que sea vulnerable a SQL injection)

	if limit == 0 {
		limit = DefaultMaxItems
	}
	_, err = o.Raw("SELECT DISTINCT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.id_facultad facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion, tr.nombre_tipo_resolucion tipo_resolucion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, administrativa.resolucion_estado re, administrativa.estado_resolucion e, administrativa.tipo_resolucion tr WHERE r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND e.nombre_estado IN('Aprobada','Expedida') AND tr.id_tipo_resolucion=r.id_tipo_resolucion ORDER BY id desc LIMIT ? OFFSET ?", limit, offset).QueryRows(&temp)

	if err != nil {
		return temp, err
	}
	for x, resoluciones := range temp {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		temp[x].FechaExpedicion = resoluciones.FechaExpedicion
	}
	return temp, nil
}

func GetAllExpedidasVigenciaPeriodo(vigencia int) (arregloIDs []ResolucionVinculacion) {
	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	_, err := o.Raw("SELECT DISTINCT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.id_facultad facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, administrativa.resolucion_estado re, administrativa.estado_resolucion e WHERE r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND r.vigencia = ? AND e.nombre_estado IN('Expedida') ORDER BY id desc;", vigencia).QueryRows(&temp)

	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	for x, resoluciones := range temp {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		temp[x].FechaExpedicion = resoluciones.FechaExpedicion
	}

	return temp
}

func GetAllExpedidasVigenciaPeriodoVinculacion(vigencia int) (arregloIDs []ResolucionVinculacion) {
	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	_, err := o.Raw("SELECT DISTINCT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.id_facultad facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, administrativa.resolucion_estado re, administrativa.estado_resolucion e WHERE r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND r.vigencia = ? AND e.nombre_estado IN('Expedida') AND id_tipo_resolucion=1 ORDER BY id desc;", vigencia).QueryRows(&temp)

	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	for x, resoluciones := range temp {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		temp[x].FechaExpedicion = resoluciones.FechaExpedicion
	}

	return temp
}
