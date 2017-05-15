package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type TrNecesidad struct {
	Necesidad                    *Necesidad
	Ffapropiacion                []*FuenteFinanciacionRubroNecesidad
	MarcoLegalNecesidad          []*MarcoLegalNecesidad
	SupervisorSolicitudNecesidad *SupervisorSolicitudNecesidad
	ActividadEconomicaNecesidad  []*ActividadEconomicaNecesidad
	Especificacion               []*TrEspecificacion
	ActividadEspecifica          []*ActividadEspecifica
	DependenciaNecesidad         *DependenciaNecesidad
	ServicioNecesidad            *ServicioNecesidad
}

type TrEspecificacion struct {
	EspecificacionTecnica *EspecificacionTecnica
	RequisitoMinimo       []*RequisitoMinimo
}

//funcion para la transaccion de solicitudes
func AddTrNecesidad(m *TrNecesidad) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	var id int64
	m.Necesidad.FechaSolicitud = time.Now()
	m.Necesidad.Numero = 0
	m.Necesidad.OtroSi = 0
	m.Necesidad.Vigencia = float64((m.Necesidad.FechaSolicitud).Year())
	m.SupervisorSolicitudNecesidad.FechaAsginacion = time.Now()
	m.Necesidad.FechaModificacion = time.Now()
	var a []int
	_, err = o.Raw("SELECT MAX(numero_elaboracion)+1 FROM administrativa.necesidad WHERE vigencia=" + strconv.Itoa((m.Necesidad.FechaSolicitud).Year()) + ";").QueryRows(&a)
	m.Necesidad.NumeroElaboracion = a[0]
	if id, err = o.Insert(m.Necesidad); err == nil {
		//m.Necesidad.Id = int(id)
		fmt.Println("Fuentes", m.Ffapropiacion)
		for _, v := range m.Ffapropiacion {
			v.SolicitudNecesidad = &Necesidad{Id: int(id)}
			//---
			if _, err = o.Insert(v); err != nil {
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Error: ¡Ocurrió un error al insertar las fuentes de financiamiento!")
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
				alerta[0] = "error"
				alerta = append(alerta, "Error: ¡Ocurrió un error al insertar los marcos legales!")
				return
			}
		}
		m.SupervisorSolicitudNecesidad.SolicitudNecesidad = &Necesidad{Id: int(id)}
		if _, err = o.Insert(m.SupervisorSolicitudNecesidad); err != nil {
			o.Rollback()
			alerta[0] = "error"
			alerta = append(alerta, "Error: ¡Ocurrió un error al insertar el supervisor de la necesidad!")
			return
		}




		m.DependenciaNecesidad.Necesidad = &Necesidad{Id: int(id)}
		if _, err = o.Insert(m.DependenciaNecesidad); err != nil {
			o.Rollback()
			alerta[0] = "error"
			alerta = append(alerta, "Error: ¡Ocurrió un error al insertar los datos de las dependencias y responsables!")
			return
		}
		if m.Necesidad.Servicio.Id==1{

					for _, ve := range m.Especificacion {
						ve.EspecificacionTecnica.SolicitudNecesidad = &Necesidad{Id: int(id)}
						//---
						if _, err = o.Insert(ve.EspecificacionTecnica); err != nil {
							o.Rollback()
							alerta[0] = "error"
							alerta = append(alerta, "Error: ¡Ocurrió un error al insertar las especificaciones técnicas!")
							return
						} else {
							for _, vr := range ve.RequisitoMinimo {
								vr.EspecificacionTecnica = ve.EspecificacionTecnica
								//---
								if _, err = o.Insert(vr); err != nil {
									o.Rollback()
									alerta[0] = "error"
									alerta = append(alerta, "Error: ¡Ocurrió un error al insertar los requisitos mínimos!")
									return
								}
							}
						}
					}
		}
		if m.Necesidad.Servicio.Id==2{
			for _, va := range m.ActividadEconomicaNecesidad {
				va.SolicitudNecesidad = &Necesidad{Id: int(id)}
				//---
				if _, err = o.Insert(va); err != nil {
					o.Rollback()
					alerta[0] = "error"
					alerta = append(alerta, "Error: ¡Ocurrió un error al insertar las actividades económicas!")
					return
				}
			}
			for _, vp := range m.ActividadEspecifica {
				vp.Necesidad = &Necesidad{Id: int(id)}
				//---
				if _, err = o.Insert(vp); err != nil {
					o.Rollback()
					alerta[0] = "error"
					alerta = append(alerta, "Error: ¡Ocurrió un error al insertar las actividades específicas!")
					return
				}
			}
			m.ServicioNecesidad.Necesidad = &Necesidad{Id: int(id)}
			if _, err = o.Insert(m.ServicioNecesidad); err != nil {
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Error: ¡Ocurrió un error al insertar servicio necesidad!")
				return
			}
		}
		o.Commit()
		alerta = append(alerta, "La solicitud con No. de elaboración "+strconv.Itoa(m.Necesidad.NumeroElaboracion)+" del "+strconv.Itoa((m.Necesidad.FechaSolicitud).Year())+" fue creada exitosamente")
		return
	} else {
		o.Rollback()
		alerta[0] = "error"
		alerta = append(alerta, "Error: ¡Ocurrió un error al insertar la necesidad!")
		return
	}
	return alerta, err
}
