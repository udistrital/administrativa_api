package models

import (
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

	if id, err = o.Insert(m.Necesidad); err != nil {
		o.Rollback()
		return
	}

	for _, v := range m.Ffapropiacion {
		v.Necesidad = &Necesidad{Id: int(id)}
		//---
		if id, err = o.Insert(v); err != nil {
			o.Rollback()
			return
		}
	}

	for idx, vm := range m.MarcoLegalNecesidad {
		vm.Necesidad = &Necesidad{Id: int(idx)}
		//---
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
			ve.EspecificacionTecnica.Necesidad = &Necesidad{Id: int(id)}
			//---
			if id, err = o.Insert(ve.EspecificacionTecnica); err != nil {
				o.Rollback()
				return
			} else {
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
	}
	if m.Necesidad.TipoContratoNecesidad.Id == 2 {
		for _, va := range m.ActividadEconomicaNecesidad {
			va.Necesidad = &Necesidad{Id: int(id)}
			//---
			if id, err = o.Insert(va); err != nil {
				o.Rollback()
				return
			}
		}
		for _, vp := range m.ActividadEspecifica {
			vp.Necesidad = &Necesidad{Id: int(id)}
			//---
			if id, err = o.Insert(vp); err != nil {
				o.Rollback()
				return
			}
		}
		m.DetalleServicioNecesidad.Necesidad = &Necesidad{Id: int(id)}
		if id, err = o.Insert(m.DetalleServicioNecesidad); err != nil {
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

//funcion para la transaccion de solicitudes
func AddTrNecesidad(m *TrNecesidad) (alertas []Alert, err error) {
	return TrNecesidadFunc(m)
}

//funcion para la transaccion de solicitudes
func UpdateTrNecesidadById(m *TrNecesidad) (alertas []Alert, err error) {
	return TrNecesidadFunc(m)
}
