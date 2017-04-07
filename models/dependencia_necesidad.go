package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type DependenciaNecesidad struct {
	Id                         int `orm:"column(id);pk;auto"`
	JefeDependenciaSolicitante int `orm:"column(jefe_dependencia_solicitante)"`
	DependenciaSolicitante     int `orm:"column(dependencia_solicitante)"`
	JefeDependenciaDestino     int `orm:"column(jefe_dependencia_destino)"`
	DependenciaDestino         int `orm:"column(dependencia_destino)"`
	Necesidad                  *Necesidad `orm:"column(necesidad);rel(fk)"`
	OrdenadorGasto             int `orm:"column(ordenador_gasto)"`
}

func (t *DependenciaNecesidad) TableName() string {
	return "dependencia_necesidad"
}

func init() {
	orm.RegisterModel(new(DependenciaNecesidad))
}

// AddDependenciaNecesidad insert a new DependenciaNecesidad into database and returns
// last inserted Id on success.
func AddDependenciaNecesidad(m *DependenciaNecesidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDependenciaNecesidadById retrieves DependenciaNecesidad by Id. Returns error if
// Id doesn't exist
func GetDependenciaNecesidadById(id int) (v *DependenciaNecesidad, err error) {
	o := orm.NewOrm()
	v = &DependenciaNecesidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDependenciaNecesidad retrieves all DependenciaNecesidad matches certain condition. Returns empty list if
// no records exist
func GetAllDependenciaNecesidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DependenciaNecesidad))
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

	var l []DependenciaNecesidad
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

// UpdateDependenciaNecesidad updates DependenciaNecesidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateDependenciaNecesidadById(m *DependenciaNecesidad) (err error) {
	o := orm.NewOrm()
	v := DependenciaNecesidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDependenciaNecesidad deletes DependenciaNecesidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDependenciaNecesidad(id int) (err error) {
	o := orm.NewOrm()
	v := DependenciaNecesidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DependenciaNecesidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
