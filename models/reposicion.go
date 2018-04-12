package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Reposicion struct {
	Id            int         `orm:"column(id);pk;auto"`
	FechaRegistro time.Time   `orm:"column(fecha_registro);type(date);auto_now"`
	Elemento      string      `orm:"column(elemento)"`
	Marca         string      `orm:"column(marca);null"`
	Serie         string      `orm:"column(serie);null"`
	Valor         int         `orm:"column(valor)"`
	Iva           int         `orm:"column(iva)"`
	Movimiento    *Movimiento `orm:"column(movimiento);rel(fk)"`
}

func (t *Reposicion) TableName() string {
	return "reposicion"
}

func init() {
	orm.RegisterModel(new(Reposicion))
}

// AddReposicion insert a new Reposicion into database and returns
// last inserted Id on success.
func AddReposicion(m *Reposicion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReposicionById retrieves Reposicion by Id. Returns error if
// Id doesn't exist
func GetReposicionById(id int) (v *Reposicion, err error) {
	o := orm.NewOrm()
	v = &Reposicion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReposicion retrieves all Reposicion matches certain condition. Returns empty list if
// no records exist
func GetAllReposicion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Reposicion)).RelatedSel(5)
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

	var l []Reposicion
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

// UpdateReposicion updates Reposicion by Id and returns error if
// the record to be updated doesn't exist
func UpdateReposicionById(m *Reposicion) (err error) {
	o := orm.NewOrm()
	v := Reposicion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReposicion deletes Reposicion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReposicion(id int) (err error) {
	o := orm.NewOrm()
	v := Reposicion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Reposicion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
