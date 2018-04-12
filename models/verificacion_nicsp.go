package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type VerificacionNicsp struct {
	Id            int         `orm:"column(id);pk;auto"`
	Verificado    bool        `orm:"column(verificado)"`
	FechaRegistro time.Time   `orm:"column(fecha_registro);type(date);auto_now"`
	ValorMercado  float64     `orm:"column(valor_mercado)"`
	Inventario    *Inventario `orm:"column(inventario);rel(fk)"`
}

func (t *VerificacionNicsp) TableName() string {
	return "verificacion_nicsp"
}

func init() {
	orm.RegisterModel(new(VerificacionNicsp))
}

// AddVerificacionNicsp insert a new VerificacionNicsp into database and returns
// last inserted Id on success.
func AddVerificacionNicsp(m *VerificacionNicsp) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetVerificacionNicspById retrieves VerificacionNicsp by Id. Returns error if
// Id doesn't exist
func GetVerificacionNicspById(id int) (v *VerificacionNicsp, err error) {
	o := orm.NewOrm()
	v = &VerificacionNicsp{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllVerificacionNicsp retrieves all VerificacionNicsp matches certain condition. Returns empty list if
// no records exist
func GetAllVerificacionNicsp(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(VerificacionNicsp)).RelatedSel(5)
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

	var l []VerificacionNicsp
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

// UpdateVerificacionNicsp updates VerificacionNicsp by Id and returns error if
// the record to be updated doesn't exist
func UpdateVerificacionNicspById(m *VerificacionNicsp) (err error) {
	o := orm.NewOrm()
	v := VerificacionNicsp{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteVerificacionNicsp deletes VerificacionNicsp by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVerificacionNicsp(id int) (err error) {
	o := orm.NewOrm()
	v := VerificacionNicsp{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&VerificacionNicsp{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
