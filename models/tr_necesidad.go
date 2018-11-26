package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

//TrNecesidad is the model for the tr-necesidad artificial model data
type TrNecesidad struct {
	Necesidad                   *Necesidad
	Ffapropiacion               []*FuenteFinanciacionRubroNecesidad
	MarcoLegalNecesidad         []*MarcoLegalNecesidad
	ActividadEconomicaNecesidad []*ActividadEconomicaNecesidad
	Especificacion              []*TrEspecificacion
	ActividadEspecifica         []*ActividadEspecifica
	DependenciaNecesidad        *DependenciaNecesidad
	DetalleServicioNecesidad    *DetalleServicioNecesidad
	ProductosNecesidad          []*ProductoRubroNecesidad
}

//TrEspecificacion is the model for the tr-especificacion artificial model data
type TrEspecificacion struct {
	EspecificacionTecnica *EspecificacionTecnica
	RequisitoMinimo       []*RequisitoMinimo
}

// AddTrNecesidad insert a new Necesidad and related tables into database and returns
// last inserted Id on success.
func AddTrNecesidad(m *TrNecesidad) (id int64, err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		return
	}

	//default values
	if m.Necesidad.ModalidadSeleccion == nil {
		m.Necesidad.ModalidadSeleccion = &ModalidadSeleccion{Id: 9}
	}
	if m.Necesidad.TipoContratoNecesidad == nil {
		m.Necesidad.TipoContratoNecesidad = &TipoContratoNecesidad{Id: 3}
	}

	m.Necesidad.FechaSolicitud = time.Now()
	m.Necesidad.Numero = 0
	m.Necesidad.Vigencia = float64((m.Necesidad.FechaSolicitud).Year())
	m.Necesidad.FechaModificacion = time.Now()
	var a []int
	fechaSolicitud := strconv.Itoa((m.Necesidad.FechaSolicitud).Year())
	_, err = o.Raw("SELECT COALESCE(MAX(numero_elaboracion), 0)+1 FROM administrativa.necesidad WHERE vigencia = ? ;", fechaSolicitud).QueryRows(&a)
	m.Necesidad.NumeroElaboracion = a[0]
	var idNecesidad int64
	if idNecesidad, err = o.Insert(m.Necesidad); err != nil {
		o.Rollback()
		return
	}

	for _, v := range m.Ffapropiacion {
		v.Necesidad = &Necesidad{Id: int(idNecesidad)}
		if id, err = o.Insert(v); err != nil {
			err = fmt.Errorf("%v %v", err, o.Rollback())
			return
		}
	}

	for _, p := range m.ProductosNecesidad {
		p.Necesidad = &Necesidad{Id: int(idNecesidad)}
		p.FechaRegistro = time.Now()
		if id, err = o.Insert(p); err != nil {
			err = o.Rollback()
			if err != nil {
				beego.Error(err)
			}
			return
		}
	}

	for _, vm := range m.MarcoLegalNecesidad {
		vm.Necesidad = &Necesidad{Id: int(idNecesidad)}
		if id, err = o.Insert(vm); err != nil {
			err = fmt.Errorf("%v %v", err, o.Rollback())
			return
		}
	}

	m.DependenciaNecesidad.Necesidad = &Necesidad{Id: int(idNecesidad)}
	if id, err = o.Insert(m.DependenciaNecesidad); err != nil {
		err = fmt.Errorf("%v %v", err, o.Rollback())
		return
	}
	if m.Necesidad.TipoContratoNecesidad.Id == 1 {

		for _, ve := range m.Especificacion {
			ve.EspecificacionTecnica.Necesidad = &Necesidad{Id: int(idNecesidad)}
			//---
			if id, err = o.Insert(ve.EspecificacionTecnica); err != nil {
				err = fmt.Errorf("%v %v", err, o.Rollback())
				return
			}
			for _, vr := range ve.RequisitoMinimo {
				vr.EspecificacionTecnica = ve.EspecificacionTecnica
				//---
				if id, err = o.Insert(vr); err != nil {
					err = fmt.Errorf("%v %v", err, o.Rollback())
					return
				}
			}

		}
	}
	if m.Necesidad.TipoContratoNecesidad.Id == 2 {
		for _, va := range m.ActividadEconomicaNecesidad {
			va.Necesidad = &Necesidad{Id: int(idNecesidad)}
			//---
			if id, err = o.Insert(va); err != nil {
				err = fmt.Errorf("%v %v", err, o.Rollback())
				return
			}
		}
		for _, vp := range m.ActividadEspecifica {
			vp.Necesidad = &Necesidad{Id: int(idNecesidad)}
			//---
			if id, err = o.Insert(vp); err != nil {
				err = fmt.Errorf("%v %v", err, o.Rollback())
				return
			}
		}
		m.DetalleServicioNecesidad.Necesidad = &Necesidad{Id: int(idNecesidad)}
		if id, err = o.Insert(m.DetalleServicioNecesidad); err != nil {
			err = fmt.Errorf("%v %v", err, o.Rollback())
			return
		}
	}
	if err = o.Commit(); err != nil {
		return 0, err
	}
	return idNecesidad, nil
}

// UpdateTrNecesidadByID updates TrNecesidad and related models by Id and returns error if
// the record to be updated doesn't exist
func UpdateTrNecesidadByID(m *TrNecesidad) (err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		return
	}

	//default values
	if m.Necesidad.ModalidadSeleccion == nil {
		m.Necesidad.ModalidadSeleccion = &ModalidadSeleccion{Id: 9}
	}
	if m.Necesidad.TipoContratoNecesidad == nil {
		m.Necesidad.TipoContratoNecesidad = &TipoContratoNecesidad{Id: 3}
	}

	m.Necesidad.FechaSolicitud = time.Now()
	m.Necesidad.Numero = 0
	m.Necesidad.Vigencia = float64((m.Necesidad.FechaSolicitud).Year())
	m.Necesidad.FechaModificacion = time.Now()
	var idNecesidad = m.Necesidad.Id
	var num int64

	if num, err = o.Update(m.Necesidad); err != nil {
		err = fmt.Errorf("%v %v", err, o.Rollback())
		return
	}
	fmt.Println("Number of records of Necesidad updated in database:", num)

	_, err = o.QueryTable(new(FuenteFinanciacionRubroNecesidad)).Filter("Necesidad__Id", idNecesidad).Delete()
	if err != nil {
		err = fmt.Errorf("%v %v", err, o.Rollback())
		return
	}
	for _, v := range m.Ffapropiacion {
		v.Necesidad = &Necesidad{Id: idNecesidad}
		if _, err = o.Insert(v); err != nil {
			err = fmt.Errorf("%v %v", err, o.Rollback())
			return
		}
	}

	_, err = o.QueryTable(new(ProductoRubroNecesidad)).Filter("Necesidad__Id", int(idNecesidad)).Delete()
	if err != nil {
		err = o.Rollback()
		if err != nil {
			beego.Error(err)
		}
		return
	}
	for _, p := range m.ProductosNecesidad {
		p.Necesidad = &Necesidad{Id: int(idNecesidad)}
		p.FechaRegistro = time.Now()
		if _, err = o.Insert(p); err != nil {
			err = o.Rollback()
			if err != nil {
				beego.Error(err)
			}
			return
		}
	}

	_, err = o.QueryTable(new(MarcoLegalNecesidad)).Filter("Necesidad__Id", idNecesidad).Delete()
	if err != nil {
		err = fmt.Errorf("%v %v", err, o.Rollback())
		return
	}
	for _, vm := range m.MarcoLegalNecesidad {
		vm.Necesidad = &Necesidad{Id: idNecesidad}
		if _, err = o.Insert(vm); err != nil {
			err = fmt.Errorf("%v %v", err, o.Rollback())
			return
		}
	}

	m.DependenciaNecesidad.Necesidad = &Necesidad{Id: idNecesidad}
	if _, err = o.Update(m.DependenciaNecesidad); err != nil {
		err = fmt.Errorf("%v %v", err, o.Rollback())
		return
	}

	if m.Necesidad.TipoContratoNecesidad.Id == 1 {
		for _, ve := range m.Especificacion {

			ve.EspecificacionTecnica.Necesidad = &Necesidad{Id: idNecesidad}
			if _, err = o.Update(ve.EspecificacionTecnica); err != nil {
				err = fmt.Errorf("%v %v", err, o.Rollback())
				return
			}
			for _, vr := range ve.RequisitoMinimo {
				vr.EspecificacionTecnica = ve.EspecificacionTecnica
				//---
				if _, err = o.Update(vr); err != nil {
					err = fmt.Errorf("%v %v", err, o.Rollback())
					return
				}
			}

		}
	}
	if m.Necesidad.TipoContratoNecesidad.Id == 2 {
		_, err = o.QueryTable(new(ActividadEconomicaNecesidad)).Filter("Necesidad__Id", idNecesidad).Delete()
		if err != nil {
			err = fmt.Errorf("%v %v", err, o.Rollback())
			return
		}
		for _, va := range m.ActividadEconomicaNecesidad {
			va.Necesidad = &Necesidad{Id: int(idNecesidad)}
			//---
			if _, err = o.InsertOrUpdate(va, "Id"); err != nil {
				err = fmt.Errorf("%v %v", err, o.Rollback())
				return
			}
		}
		_, err = o.QueryTable(new(ActividadEspecifica)).Filter("Necesidad__Id", idNecesidad).Delete()
		if err != nil {
			err = fmt.Errorf("%v %v", err, o.Rollback())
			return
		}
		for _, vp := range m.ActividadEspecifica {
			vp.Necesidad = &Necesidad{Id: idNecesidad}
			//---
			if _, err = o.InsertOrUpdate(vp, "Id"); err != nil {
				err = fmt.Errorf("%v %v", err, o.Rollback())
				return
			}
		}
		dsn := new(DetalleServicioNecesidad)
		m.DetalleServicioNecesidad.Necesidad = &Necesidad{Id: idNecesidad}
		if err = o.QueryTable(new(DetalleServicioNecesidad)).Filter("Necesidad", idNecesidad).One(dsn); err == nil {
			m.DetalleServicioNecesidad.Id = dsn.Id
			if _, err = o.Update(m.DetalleServicioNecesidad); err != nil {
				err = fmt.Errorf("%v %v", err, o.Rollback())
				return
			}
		} else {
			m.DependenciaNecesidad.Id = 0
			if _, err = o.Insert(m.DetalleServicioNecesidad); err != nil {
				err = fmt.Errorf("%v %v", err, o.Rollback())
				return
			}
		}
	}
	if err = o.Commit(); err != nil {
		return err
	}
	return
}
