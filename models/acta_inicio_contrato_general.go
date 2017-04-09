package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	//"strconv"
	"github.com/astaxie/beego/orm"
)

type ActaInicioContratoGeneral struct {
	Id             int       `orm:"column(id);pk"`
	NumeroContrato *ContratoGeneral    `orm:"rel(one);column(numero_contrato);null"`
	Vigencia       int       `orm:"column(vigencia);null"`
	Vigencia_contrato       int       `orm:"column(vigencia_contrato);null"`
	FechaInicio    time.Time `orm:"column(fecha_inicio);type(date);null"`
	FechaFin       time.Time `orm:"column(fecha_fin);type(date);null"`
	Descripcion    string    `orm:"column(descripcion);null"`
	Usuario        string    `orm:"column(usuario);null"`
	ObjetoContrato               string           `orm:"column(objeto_contrato);null"`
	PlazoEjecucion               int              `orm:"column(plazo_ejecucion)"`
//	FormaPago                    *Parametros      `orm:"column(forma_pago);rel(fk)"`
	//OrdenadorGasto               *ArgoOrdenadores `orm:"column(ordenador_gasto);rel(fk)"`
	ClausulaRegistroPresupuestal bool             `orm:"column(clausula_registro_presupuestal);null"`
	SedeSolicitante              string           `orm:"column(sede_solicitante);null"`
	DependenciaSolicitante       string           `orm:"column(dependencia_solicitante);null"`
}

func init() {
	orm.RegisterModel(new(ActaInicioContratoGeneral))
}

// AddActaInicioContratoGeneral insert a new ActaInicioContratoGeneral into database and returns
// last inserted Id on success.
func AddActaInicioContratoGeneral(m *ActaInicioContratoGeneral) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetActaInicioContratoGeneralById retrieves ActaInicioContratoGeneral by Id. Returns error if
// Id doesn't exist
func GetActaInicioContratoGeneralById(id int) (v *ActaInicioContratoGeneral, err error) {
	o := orm.NewOrm()
	v = &ActaInicioContratoGeneral{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllActaInicioContratoGeneral retrieves all ActaInicioContratoGeneral matches certain condition. Returns empty list if
// no records exist
func GetAllActaInicioContratoGeneral(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ActaInicioContratoGeneral))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ActaInicioContratoGeneral
	qs = qs.OrderBy(sortFields...).RelatedSel(5)

	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateActaInicioContratoGeneral updates ActaInicioContratoGeneral by Id and returns error if
// the record to be updated doesn't exist
func UpdateActaInicioContratoGeneralById(m *ActaInicioContratoGeneral) (err error) {
	o := orm.NewOrm()
	v := ActaInicioContratoGeneral{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteActaInicioContratoGeneral deletes ActaInicioContratoGeneral by Id and returns error if
// the record to be deleted doesn't exist
func DeleteActaInicioContratoGeneral(id int) (err error) {
	o := orm.NewOrm()
	v := ActaInicioContratoGeneral{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ActaInicioContratoGeneral{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func JoinActaInicioContratoGeneral(vigencia int)(Acta_Inicio_Contrato_General []ActaInicioContratoGeneral){
	o := orm.NewOrm()
	var temp []ActaInicioContratoGeneral
	fmt.Println(vigencia)
	//vigenciaString:= strconv.Itoa(vigencia)
	_, err := o.Raw("SELECT acta_inicio.id as id,acta_inicio.numero_contrato as numero_contrato,acta_inicio.vigencia as vigencia,contrato_general.vigencia as vigencia_contrato, acta_inicio.fecha_inicio as fecha_inicio,acta_inicio.descripcion,contrato_general.objeto_contrato as objeto_contrato,contrato_general.plazo_ejecucion as plazo_ejecucion,contrato_general.forma_pago as forma_pago, contrato_general.ordenador_gasto as ordenador_gasto, contrato_general.clausula_registro_presupuestal as clausula_registro_presupuestal, contrato_general.sede_solicitante as sede_solicitante, contrato_general.dependencia_solicitante as dependencia_solicitante FROM argo.acta_inicio INNER JOIN argo.contrato_general ON acta_inicio.numero_contrato = contrato_general.numero_contrato and acta_inicio.vigencia = contrato_general.vigencia").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
		return temp
}
