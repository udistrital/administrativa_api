package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type TrNecesidad struct {
	Necesidad                   *Necesidad
	Ffapropiacion               []*FuenteFinanciacionRubroNecesidad
	MarcoLegalNecesidad         []*MarcoLegalNecesidad
	ActividadEconomicaNecesidad []*ActividadEconomicaNecesidad
	Especificacion              []*TrEspecificacion
	ActividadEspecifica         []*ActividadEspecifica
	DependenciaNecesidad        *DependenciaNecesidad
	DetalleServicioNecesidad    *DetalleServicioNecesidad
}

type TrEspecificacion struct {
	EspecificacionTecnica *EspecificacionTecnica
	RequisitoMinimo       []*RequisitoMinimo
}

func AddTrNecesidad(m *TrNecesidad) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()

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
	_, err = o.Raw("SELECT COALESCE(MAX(numero_elaboracion), 0)+1 FROM administrativa.necesidad WHERE vigencia=" + strconv.Itoa((m.Necesidad.FechaSolicitud).Year()) + ";").QueryRows(&a)
	m.Necesidad.NumeroElaboracion = a[0]
	var idNecesidad int64
	if idNecesidad, err = o.Insert(m.Necesidad); err != nil {
		o.Rollback()
		return
	}

	for _, v := range m.Ffapropiacion {
		v.Necesidad = &Necesidad{Id: int(idNecesidad)}
		if id, err = o.Insert(v); err != nil {
			o.Rollback()
			return
		}
	}

	for _, vm := range m.MarcoLegalNecesidad {
		vm.Necesidad = &Necesidad{Id: int(idNecesidad)}
		if id, err = o.Insert(vm); err != nil {
			o.Rollback()
			return
		}
	}

	m.DependenciaNecesidad.Necesidad = &Necesidad{Id: int(id)}
	if id, err = o.Insert(m.DependenciaNecesidad); err != nil {
		o.Rollback()
		return
	}
	if m.Necesidad.TipoContratoNecesidad.Id == 1 {

		for _, ve := range m.Especificacion {
			ve.EspecificacionTecnica.Necesidad = &Necesidad{Id: int(idNecesidad)}
			//---
			if id, err = o.Insert(ve.EspecificacionTecnica); err != nil {
				o.Rollback()
				return
			}
			for _, vr := range ve.RequisitoMinimo {
				vr.EspecificacionTecnica = ve.EspecificacionTecnica
				//---
				if id, err = o.Insert(vr); err != nil {
					o.Rollback()
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
				o.Rollback()
				return
			}
		}
		for _, vp := range m.ActividadEspecifica {
			vp.Necesidad = &Necesidad{Id: int(idNecesidad)}
			//---
			if id, err = o.Insert(vp); err != nil {
				o.Rollback()
				return
			}
		}
		m.DetalleServicioNecesidad.Necesidad = &Necesidad{Id: int(idNecesidad)}
		if id, err = o.Insert(m.DetalleServicioNecesidad); err != nil {
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

func UpdateTrNecesidadById(m *TrNecesidad) (err error) {
	o := orm.NewOrm()
	o.Begin()

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
	_, err = o.Raw("SELECT COALESCE(MAX(numero_elaboracion), 0)+1 FROM administrativa.necesidad WHERE vigencia=" + strconv.Itoa((m.Necesidad.FechaSolicitud).Year()) + ";").QueryRows(&a)
	if err != nil {
		o.Rollback()
		return
	}
	m.Necesidad.NumeroElaboracion = a[0]
	var idNecesidad int = m.Necesidad.Id
	var num int64

	if num, err = o.Update(m.Necesidad); err != nil {
		o.Rollback()
		return
	}
	fmt.Println("Number of records of Necesidad updated in database:", num)

	var Ffapropiaciones []FuenteFinanciacionRubroNecesidad
	num, err = o.QueryTable(m.Ffapropiacion[0]).Filter("Necesidad__Id", int(idNecesidad)).All(&Ffapropiaciones)
	if err != nil {
		o.Rollback()
		return
	}
	for _, v1 := range Ffapropiaciones {
		exist := false
		for _, v2 := range m.Ffapropiacion {
			v2.Necesidad = &Necesidad{Id: int(idNecesidad)}
			if _, err = o.InsertOrUpdate(v2, "Id"); err != nil {
				o.Rollback()
				return
			}
			if v1.Id == v2.Id {
				exist = true
				break
			}
		}
		if !exist {
			if _, err = o.Delete(&v1); err != nil {
				o.Rollback()
				return
			}
		}
	}

	var MarcosLegales []MarcoLegalNecesidad
	num, err = o.QueryTable(m.MarcoLegalNecesidad[0]).Filter("Necesidad__Id", idNecesidad).All(&MarcosLegales)

	for _, v1 := range MarcosLegales {
		exist := false
		for _, v2 := range m.MarcoLegalNecesidad {
			v2.Necesidad = &Necesidad{Id: int(idNecesidad)}
			o.InsertOrUpdate(v2, "Id")
			if v1.Id == v2.Id {
				exist = true
				break
			}
		}
		if !exist {
			if _, err = o.Delete(&v1); err != nil {
				o.Rollback()
				return
			}
		}
	}

	m.DependenciaNecesidad.Necesidad = &Necesidad{Id: int(idNecesidad)}
	if _, err = o.InsertOrUpdate(m.DependenciaNecesidad, "Id"); err != nil {
		o.Rollback()
		return
	}

	if m.Necesidad.TipoContratoNecesidad.Id == 1 {
		for _, ve := range m.Especificacion {

			ve.EspecificacionTecnica.Necesidad = &Necesidad{Id: int(idNecesidad)}
			if _, err = o.InsertOrUpdate(ve.EspecificacionTecnica, "Id"); err != nil {
				o.Rollback()
				return
			}
			for _, vr := range ve.RequisitoMinimo {
				vr.EspecificacionTecnica = ve.EspecificacionTecnica
				//---
				if _, err = o.InsertOrUpdate(vr, "Id"); err != nil {
					o.Rollback()
					return
				}
			}

		}
	}
	if m.Necesidad.TipoContratoNecesidad.Id == 2 {
		for _, va := range m.ActividadEconomicaNecesidad {
			va.Necesidad = &Necesidad{Id: int(idNecesidad)}
			//---
			if _, err = o.InsertOrUpdate(va, "Id"); err != nil {
				o.Rollback()
				return
			}
		}
		for _, vp := range m.ActividadEspecifica {
			vp.Necesidad = &Necesidad{Id: int(idNecesidad)}
			//---
			if _, err = o.InsertOrUpdate(vp, "Id"); err != nil {
				o.Rollback()
				return
			}
		}
		m.DetalleServicioNecesidad.Necesidad = &Necesidad{Id: int(idNecesidad)}
		if _, err = o.InsertOrUpdate(m.DetalleServicioNecesidad, "Id"); err != nil {
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}
