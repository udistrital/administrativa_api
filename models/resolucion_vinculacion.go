package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

type ResolucionVinculacion struct {
	Id                 int       `orm:"column(id);pk;auto"`
	Estado             string    `orm:"column(estado)"`
	Numero             string    `orm:"column(numero)"`
	Vigencia           int       `orm:"column(vigencia)"`
	Facultad           int       `orm:"column(facultad)"`
	NivelAcademico     string    `orm:"column(nivel_academico)"`
	Dedicacion         string    `orm:"column(dedicacion)"`
	FechaExpedicion    time.Time `orm:"column(fecha_expedicion);type(timestamp without time zone)"`
	NumeroSemanas      int       `orm:"column(numero_semanas)"`
	Periodo            int       `orm:"column(periodo)"`
	TipoResolucion     string    `orm:"column(tipo_resolucion)"`
	IdDependenciaFirma int       `orm:"column(dependencia_firma)"`
}

var (
	columnNames = make(map[string]string)
)

func init() {
	orm.RegisterModel(new(ResolucionVinculacion))
	t := reflect.TypeOf(ResolucionVinculacion{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("orm")
		column := ""
		_, err := fmt.Sscanf(strings.Split(tag, ";")[0], "column(%s)", &column)
		if err != nil {
			beego.Error(err)
		}
		columnNames[field.Name] = column[:len(column)-1]
	}
}

func GetAllResolucionVinculacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []ResolucionVinculacion, err error) {
	o := orm.NewOrm()

	if limit == 0 {
		limit = DefaultMaxItems
	}
	qs := make([]Operation, 0)
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = append(qs, filter(k, (v == "true" || v == "1")))
		} else {
			qs = append(qs, filter(k, v))
		}
	}

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return
	}

	// _, err = o.Raw(`
	qb.Select(
		"DISTINCT r.id_resolucion id",
		"e.nombre_estado estado",
		"r.numero_resolucion numero",
		"r.vigencia vigencia",
		"r.periodo periodo",
		"rv.id_facultad facultad",
		"rv.nivel_academico nivel_academico",
		"rv.dedicacion dedicacion",
		"r.numero_semanas numero_semanas",
		"r.fecha_expedicion fecha_expedicion",
		"tipo.nombre_tipo_resolucion tipo_resolucion",
		"r.id_dependencia_firma dependencia_firma").
		From(
			"administrativa.resolucion r",
			"administrativa.resolucion_vinculacion_docente rv",
			"administrativa.resolucion_estado re",
			"administrativa.estado_resolucion e",
			"administrativa.tipo_resolucion tipo").
		Where("r.id_resolucion=rv.id_resolucion").
		And("re.resolucion=r.id_resolucion").
		And("r.id_tipo_resolucion=tipo.id_tipo_resolucion").
		And("re.estado=e.id").
		And("re.estado!=6").
		And("re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion)")

	qb2, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return ml, err
	}
	qb2.Select("*").
		From(qb.Subquery(qb.String(), "T"))

	// query externo
	flag := true
	for _, v := range qs {
		columnName, ok := columnNames[v.Field]
		if !ok {
			return ml, fmt.Errorf("inexistent fielq in query")
		}
		tmp := fmt.Sprintf("T.%s::VARCHAR %s", columnName, v.Op)
		if flag {
			qb2.Where(tmp)
			flag = false
		} else {
			qb2.And(tmp)
		}
	}

	qb2.OrderBy("id").
		Desc().
		Limit(int(limit)).
		Offset(int(offset))
	_, err = o.Raw(qb2.String()).QueryRows(&ml)
	if err != nil {
		return ml, err
	}
	for x, resoluciones := range ml {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		ml[x].FechaExpedicion = resoluciones.FechaExpedicion
	}
	return ml, nil
}

func GetAllResolucionAprobada(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (arregloIDs []ResolucionVinculacion, err error) {
	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	//_, err := o.Raw("SELECT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, d.nombre facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, oikos.dependencia d, administrativa.resolucion_estado re, administrativa.estado_resolucion e WHERE rv.id_facultad=d.id AND r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND r.id_tipo_resolucion=1 ORDER BY id desc;").QueryRows(&temp)
	//TODO: dar soporte a query (sin dejar que sea vulnerable a SQL injection)

	if limit == 0 {
		limit = DefaultMaxItems
	}

	qs := make([]Operation, 0)
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = append(qs, filter(k, (v == "true" || v == "1")))
		} else {
			qs = append(qs, filter(k, v))
		}
	}

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return
	}
	qb.Select(
		"DISTINCT r.id_resolucion id",
		"e.nombre_estado estado",
		"r.numero_resolucion numero",
		"r.vigencia vigencia",
		"r.periodo periodo",
		"rv.id_facultad facultad",
		"rv.nivel_academico nivel_academico",
		"rv.dedicacion dedicacion",
		"r.numero_semanas numero_semanas",
		"r.fecha_expedicion fecha_expedicion",
		"tr.nombre_tipo_resolucion tipo_resolucion",
		"r.id_dependencia_firma dependencia_firma",
	).
		From(
			"administrativa.resolucion r",
			"administrativa.resolucion_vinculacion_docente rv",
			"administrativa.resolucion_estado re",
			"administrativa.estado_resolucion e",
			"administrativa.tipo_resolucion tr",
		).
		Where("r.id_resolucion=rv.id_resolucion").
		And("re.resolucion=r.id_resolucion").
		And("re.estado=e.id").
		And("re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND e.nombre_estado IN('Aprobada','Expedida')").
		And("tr.id_tipo_resolucion=r.id_tipo_resolucion")

	qb2, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		return arregloIDs, err
	}
	qb2.Select("*").
		From(qb.Subquery(qb.String(), "T"))

	// query externo
	flag := true
	for _, v := range qs {
		columnName, ok := columnNames[v.Field]
		beego.Debug(columnName)
		if !ok {
			return arregloIDs, fmt.Errorf("inexistent fielq in query")
		}
		tmp := fmt.Sprintf("T.%s::VARCHAR %s", columnName, v.Op)
		if flag {
			qb2.Where(tmp)
			flag = false
		} else {
			qb2.And(tmp)
		}
	}

	qb2.OrderBy("id").
		Desc().
		Limit(int(limit)).
		Offset(int(offset))

	_, err = o.Raw(qb2.String()).QueryRows(&temp)
	if err != nil {
		return arregloIDs, err
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
	_, err := o.Raw("SELECT DISTINCT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, r.periodo periodo, rv.id_facultad facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.numero_semanas numero_semanas,r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, administrativa.resolucion_estado re, administrativa.estado_resolucion e WHERE r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND r.vigencia = ? AND e.nombre_estado IN('Expedida') AND id_tipo_resolucion IN (1,3,4) ORDER BY id desc;", vigencia).QueryRows(&temp)

	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	for x, resoluciones := range temp {
		resoluciones.FechaExpedicion = resoluciones.FechaExpedicion.UTC()
		temp[x].FechaExpedicion = resoluciones.FechaExpedicion
	}

	return temp
}
