package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/administrativa_crud_api/utilidades"
	"reflect"
	"strings"
	"time"
)

type SolicitudRp struct {
	Vigencia             int       `orm:"column(vigencia)"`
	FechaSolicitud       time.Time `orm:"column(fecha_solicitud);type(date);null;auto_now"`
	Cdp                  int       `orm:"column(cdp)"`
	Expedida             bool      `orm:"column(expedida);null"`
	NumeroContrato       string    `orm:"column(numero_contrato)"`
	VigenciaContrato     int       `orm:"column(vigencia_contrato);null"`
	Id                   int       `orm:"column(id);pk;auto"`
	TipoCompromiso       int       `orm:"column(tipo_compromiso);null"`
	JustificacionRechazo string    `orm:"column(justificacion_rechazo);null"`
	Masivo               bool      `orm:"column(masivo);null"`
	Proveedor            int       `orm:"column(id_proveedor);null"`
	NumeroCompromiso     int       `orm:"column(numero_compromiso);null"`
}

func (t *SolicitudRp) TableName() string {
	return "solicitud_rp"
}

func init() {
	orm.RegisterModel(new(SolicitudRp))
}

// AddSolicitudRp insert a new SolicitudRp into database and returns
// last inserted Id on success.
func AddSolicitudRp(m *SolicitudRp) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// AddSolicitudRp insert a new SolicitudRp into database and returns
// last inserted Id on success.
func AddSolicitudRpTr(m []map[string]interface{}) (res []map[string]interface{}) {
	o := orm.NewOrm()
	solicitud := SolicitudRp{}
	var rubros []DisponibilidadApropiacionSolicitudRp
	o.Begin()
	for _, data := range m {
		if e := utilidades.FillStruct(data["solicitudRp"], &solicitud); e == nil {
			fmt.Println(solicitud)
			solicitud.Vigencia = time.Now().Year()
			if id, err := o.Insert(&solicitud); err == nil {
				if e := utilidades.FillStruct(data["rubros"], &rubros); e == nil {
					for _, row := range rubros {
						solicitud.Id = int(id)
						row.SolicitudRp = &solicitud
						if _, err = o.Insert(&row); err != nil {
							o.Rollback()
							alrt := map[string]interface{}{"Code": "E_SRP001", "Type": "error", "Body": solicitud}
							res = append(res, alrt)
							return
						}
					}
				} else {
					fmt.Println("error conversion rubros")
					o.Rollback()
					alrt := map[string]interface{}{"Code": "E_SRP002", "Type": "error", "Body": e.Error()}
					res = append(res, alrt)
					return
				}
			} else {
				fmt.Println("error insersion solicitud")
				o.Rollback()
				alrt := map[string]interface{}{"Code": "E_SRP002", "Type": "error", "Body": err.Error()}
				res = append(res, alrt)
				return
			}
		} else {
			fmt.Println("error conversion")
			o.Rollback()
			alrt := map[string]interface{}{"Code": "E_SRP002", "Type": "error", "Body": e.Error()}
			res = append(res, alrt)
			return
		}
		alrt := map[string]interface{}{"Code": "S_SRP001", "Type": "success", "Body": solicitud}
		res = append(res, alrt)
		solicitud = SolicitudRp{}
	}
	o.Commit()
	return
}

// GetSolicitudRpById retrieves SolicitudRp by Id. Returns error if
// Id doesn't exist
func GetSolicitudRpById(id int) (v *SolicitudRp, err error) {
	o := orm.NewOrm()
	v = &SolicitudRp{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudRp retrieves all SolicitudRp matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudRp(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudRp))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
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

	var l []SolicitudRp
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

// UpdateSolicitudRp updates SolicitudRp by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudRpById(m *SolicitudRp) (err error) {
	o := orm.NewOrm()
	v := SolicitudRp{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudRp deletes SolicitudRp by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudRp(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudRp{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudRp{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
