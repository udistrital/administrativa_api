package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type SolicitudDisponibilidad struct {
	Id             int        `orm:"column(id);pk;auto"`
	Numero         int        `orm:"column(numero)"`
	Vigencia       float64    `orm:"column(vigencia)"`
	FechaSolicitud time.Time  `orm:"column(fecha_solicitud);type(date)"`
	Necesidad      *Necesidad `orm:"column(necesidad);rel(fk)"`
	Expedida       bool       `orm:"column(expedida)"`
}

func (t *SolicitudDisponibilidad) TableName() string {
	return "solicitud_disponibilidad"
}

func init() {
	orm.RegisterModel(new(SolicitudDisponibilidad))
}

// AddSolicitudDisponibilidad insert a new SolicitudDisponibilidad into database and returns
// last inserted Id on success.
func AddSolicitudDisponibilidad(m *SolicitudDisponibilidad) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	m.FechaSolicitud = time.Now()
	m.Vigencia = float64((m.FechaSolicitud).Year())
	m.Expedida = false
	var a []int
	_,err = o.Raw("SELECT MAX(numero)+1 FROM administrativa.solicitud_disponibilidad WHERE vigencia="+strconv.Itoa((m.FechaSolicitud).Year())+";").QueryRows(&a)
	m.Numero = a[0]
	if _, err = o.Insert(m); err!=nil{
		alerta[0] = "error"
		alerta = append(alerta, "Error: ¡Ocurrió un error al insertar la solicitud de disponibilidad!")
		o.Rollback()
		return
	} else {
		m.Necesidad.Estado.Id=7;
		if _,err = o.Update(m.Necesidad); err!=nil{
			alerta[0] = "error"
			alerta = append(alerta, "Error: ¡Ocurrió un error al actualizar el estado de la necesidad!")
			o.Rollback()
			return
		}
	}
	alerta = append(alerta, "La solicitud de disponibilidad No. "+strconv.Itoa(m.Numero)+" del "+strconv.Itoa((m.FechaSolicitud).Year())+" fué creada exitosamente")
	o.Commit()
	return alerta, err
}

// GetSolicitudDisponibilidadById retrieves SolicitudDisponibilidad by Id. Returns error if
// Id doesn't exist
func GetSolicitudDisponibilidadById(id int) (v *SolicitudDisponibilidad, err error) {
	o := orm.NewOrm()
	v = &SolicitudDisponibilidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudDisponibilidad retrieves all SolicitudDisponibilidad matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudDisponibilidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudDisponibilidad)).RelatedSel(5)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []SolicitudDisponibilidad
	qs = qs.OrderBy(sortFields...)
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

// UpdateSolicitudDisponibilidad updates SolicitudDisponibilidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudDisponibilidadById(m *SolicitudDisponibilidad) (err error) {
	o := orm.NewOrm()
	v := SolicitudDisponibilidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudDisponibilidad deletes SolicitudDisponibilidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudDisponibilidad(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudDisponibilidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudDisponibilidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
