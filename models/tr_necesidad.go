package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
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

//funcion para la transaccion de solicitudes
func AddTrNecesidad(m *TrNecesidad) (alertas []Alert, err error) {
	o := orm.NewOrm()
	o.Begin()
	var id int64
	m.Necesidad.FechaSolicitud = time.Now()
	m.Necesidad.Numero = 0
	m.Necesidad.Vigencia = float64((m.Necesidad.FechaSolicitud).Year())
	m.Necesidad.FechaModificacion = time.Now()
	var a []int
	_, err = o.Raw("SELECT COALESCE(MAX(numero_elaboracion), 0)+1 FROM administrativa.necesidad WHERE vigencia=" + strconv.Itoa((m.Necesidad.FechaSolicitud).Year()) + ";").QueryRows(&a)
	m.Necesidad.NumeroElaboracion = a[0]
	if id, err = o.Insert(m.Necesidad); err == nil {
		//m.Necesidad.Id = int(id)
		fmt.Println("Fuentes", m.Ffapropiacion)
		for _, v := range m.Ffapropiacion {
			v.Necesidad = &Necesidad{Id: int(id)}
			//---
			if _, err = o.Insert(v); err != nil {
				o.Rollback()
				beego.Error(err)
				alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar las fuentes de financiamiento!"})
				return
			}
		}
		//--
		fmt.Println("Marco", m.MarcoLegalNecesidad)
		for _, vm := range m.MarcoLegalNecesidad {
			vm.Necesidad = &Necesidad{Id: int(id)}
			//---
			if _, err = o.Insert(vm); err != nil {
				o.Rollback()
				beego.Error(err)
				alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar los marcos legales!"})
				return
			}
		}

		m.DependenciaNecesidad.Necesidad = &Necesidad{Id: int(id)}
		if _, err = o.Insert(m.DependenciaNecesidad); err != nil {
			o.Rollback()
			beego.Error(err)
			alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar los datos de las dependencias y responsables!"})
			return
		}
		if m.Necesidad.TipoContratoNecesidad.Id == 1 {

			for _, ve := range m.Especificacion {
				ve.EspecificacionTecnica.Necesidad = &Necesidad{Id: int(id)}
				//---
				if _, err = o.Insert(ve.EspecificacionTecnica); err != nil {
					o.Rollback()
					beego.Error(err)
					alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar las especificaciones técnicas!"})
					return
				} else {
					for _, vr := range ve.RequisitoMinimo {
						vr.EspecificacionTecnica = ve.EspecificacionTecnica
						//---
						if _, err = o.Insert(vr); err != nil {
							o.Rollback()
							beego.Error(err)
							alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar los requisitos mínimos!"})
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
				if _, err = o.Insert(va); err != nil {
					o.Rollback()
					beego.Error(err)
					alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar las actividades económicas!"})
					return
				}
			}
			for _, vp := range m.ActividadEspecifica {
				vp.Necesidad = &Necesidad{Id: int(id)}
				//---
				if _, err = o.Insert(vp); err != nil {
					o.Rollback()
					beego.Error(err)
					alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar las actividades específicas!"})
					return
				}
			}
			m.DetalleServicioNecesidad.Necesidad = &Necesidad{Id: int(id)}
			if _, err = o.Insert(m.DetalleServicioNecesidad); err != nil {
				o.Rollback()
				beego.Error(err)
				alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar servicio necesidad!"})
				return
			}
		}
		o.Commit()
		alertas = append(alertas, Alert{Type: AlertSucess, Body: m})

		return
	} else {
		o.Rollback()
		beego.Error(err)
		alertas = append(alertas, Alert{Type: AlertError, Body: "Error: ¡Ocurrió un error al insertar la necesidad!"})
		return
	}
	return alertas, err
}
