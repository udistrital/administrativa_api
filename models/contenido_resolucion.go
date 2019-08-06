package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Paragrafo struct {
	Id     int
	Numero int
	Texto  string
}

type Articulo struct {
	Id         int
	Numero     int
	Texto      string
	Paragrafos []Paragrafo
}

type ResolucionCompleta struct {
	Vinculacion   ResolucionVinculacionDocente
	Consideracion string
	Preambulo     string
	Vigencia      int
	Numero        string
	Id            int
	Articulos     []Articulo
	Titulo        string
}

func GetOneResolucionCompleta(idResolucion string) (resolucion ResolucionCompleta) {
	o := orm.NewOrm()
	var temp []Resolucion
	idRes, _ := strconv.Atoi(idResolucion)

	_, err := o.QueryTable("resolucion").Filter("id_resolucion", idRes).All(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	resolucionCompleta := ResolucionCompleta{Id: temp[0].Id, Consideracion: temp[0].ConsideracionResolucion, Preambulo: temp[0].PreambuloResolucion, Vigencia: temp[0].Vigencia, Numero: temp[0].NumeroResolucion, Titulo: temp[0].Titulo}

	var arts []ComponenteResolucion
	_, err2 := o.QueryTable("componente_resolucion").Filter("resolucion_id", idRes).Filter("tipo_componente", "Articulo").OrderBy("numero").All(&arts)
	if err2 == nil {
		fmt.Println("Consulta exitosa")
	}

	var articulos []Articulo

	for _, art := range arts {
		articulo := Articulo{Id: art.Id, Numero: art.Numero, Texto: art.Texto}

		var pars []ComponenteResolucion
		_, err3 := o.QueryTable("componente_resolucion").Filter("resolucion_id", idRes).Filter("tipo_componente", "Paragrafo").Filter("componente_padre", articulo.Id).OrderBy("numero").All(&pars)
		if err3 == nil {
			fmt.Println("Consulta exitosa")
		}

		var paragrafos []Paragrafo

		for _, par := range pars {
			paragrafo := Paragrafo{Id: par.Id, Numero: par.Numero, Texto: par.Texto}
			paragrafos = append(paragrafos, paragrafo)
		}

		articulo.Paragrafos = paragrafos

		articulos = append(articulos, articulo)
	}
	resolucionCompleta.Articulos = articulos
	return resolucionCompleta
}

func UpdateResolucionCompletaById(m *ResolucionCompleta) (err error) {
	o := orm.NewOrm()
	v := Resolucion{Id: m.Id}
	if err = o.Read(&v); err == nil {
		v.NumeroResolucion = m.Numero
		v.Titulo = m.Titulo
		_, err = o.Update(&v)
	} else {
		return
	}
	idResolucionStr := strconv.Itoa(m.Id)
	r := m.Vinculacion
	fmt.Println(r.Id)
	a := ResolucionVinculacionDocente{Id: r.Id}
	if err = o.Read(&a); err == nil {
		_, err = o.Update(&r)
	} else {
		return
	}
	if err = o.Read(&v); err == nil {
		v.ConsideracionResolucion = m.Consideracion
		v.PreambuloResolucion = m.Preambulo
		v.NumeroResolucion = m.Numero
		fmt.Println(v)
		if err := UpdateResolucionById(&v); err != nil {
		}

		resolucionCompleta := GetOneResolucionCompleta(idResolucionStr)

		for _, articulo := range resolucionCompleta.Articulos {
			if articulo.Paragrafos != nil {
				for _, paragrafo := range articulo.Paragrafos {
					if err := DeleteComponenteResolucion(paragrafo.Id); err != nil {
					}
				}
			}
			if err := DeleteComponenteResolucion(articulo.Id); err != nil {
			}
		}

		for indexArticulo, articulo := range m.Articulos {
			componenteArticulo := ComponenteResolucion{ResolucionId: &Resolucion{Id: m.Id}, Texto: articulo.Texto, Numero: indexArticulo + 1, TipoComponente: "Articulo"}
			if _, err := AddComponenteResolucion(&componenteArticulo); err == nil {
				if articulo.Paragrafos != nil {
					for indexParagrafo, paragrafo := range articulo.Paragrafos {
						componenteParagrafo := ComponenteResolucion{ResolucionId: &Resolucion{Id: m.Id}, Texto: paragrafo.Texto, Numero: indexParagrafo + 1, TipoComponente: "Paragrafo", ComponentePadre: &ComponenteResolucion{Id: componenteArticulo.Id}}
						if _, err := AddComponenteResolucion(&componenteParagrafo); err == nil {

						}
					}
				}
			}
		}
	}
	return
}

func GetTemplateResolucion(dedicacion, nivel, periodo, tipo string, numero string) (res ResolucionCompleta) {
	var resolucion ResolucionCompleta
	var articulos []Articulo
	var articulo Articulo
	var paragrafo Paragrafo
	var vigencia, _, _ = time.Now().Date()
	//var accion string
	var periodoStr string
	var nombreDedicacion string

	switch periodo {
	case "1":
		periodoStr = "primer"
	case "2":
		periodoStr = "segundo"
	case "3":
		periodoStr = "tercer"
	}

	switch dedicacion {
	case "HCP":
		if nivel == "PREGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), Honorarios hora catedra, en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\nQue mediante Resolución Nro. 001 del 15 de febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial, hora cátedra, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos, en los términos del artículo 73 de la Ley 30 de 1992.\n\nQue el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue mediante Resolución 003 de enero 14 de 2019 y Resolución 021 de febrero 26 del 2019, por medio de las cuales expiden el calendario académico para el año 2019 para las actividades académicas tercer periodo académico del año 2019 que corresponde del doce (12) de agosto al trece (13) de diciembre del 2019.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que realicen las actividades lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital.\n\nQue de conformidad con el artículo 2° del Decreto 1019 de seis (06) de junio de 2019 establece “A partir del 1° de enero de 2019, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en catorce mil doscientos diez pesos ($14.210) moneda corriente”.\n\nQue mediante resolución 312 del veinticinco (25) de julio del 2019 emitida por Rectoría de la Universidad Distrital Francisco José de Caldas en donde acoge y aplica, en lo pertinente única y expresamente al valor del punto para los docentes de Vinculación Especial hora cátedra, honorarios, Medio Tiempo Ocasional y Tiempo Completo Ocasional en pregrado en el valor de catorce mil doscientos diez pesos ($14.210) moneda corriente.\n\nQue en materia de Seguridad y Salud en el Trabajo SST para docentes ocasionales de la Universidad Distrital se deben informar las obligaciones específicas se encuentra contemplado en el Decreto 1072 de 2015 artículo 2.2.4.2.2.16. y la resolución de rectoría No. 624 de 2017, el docente ocasional, hora catedra y honorarios debe cumplir con las normas del sistema general de riesgos laborales.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. _______ del _______ de enero de 2019.\n\nQue, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "Vincular para el " + periodoStr + " Periodo Académico del año " + strconv.Itoa(vigencia) + " en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, como docentes en la modalidad de Hora Cátedra de Vinculación Especial, en el escalafón y dedicación establecidas en la tabla, para el periodo comprendido entre el doce (12) de agosto al trece (13) de diciembre del 2019, a los siguientes docentes:"}
		}
		if nivel == "POSGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), Honorarios hora catedra, en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\nQue mediante Resolución Nro. 001 del 15 de febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial, hora cátedra, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos, en los términos del artículo 73 de la Ley 30 de 1992.\n\nQue el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue, mediante Acuerdo 005 de julio 27 de 2001, igualmente al Acuerdo 006 de julio 19 del 2002 por medio del cual fija el valor Hora Cátedra honorarios y se establece el número máximo de horas, para los docentes que presten servicios a la Universidad Distrital en los programas de Postgrado.\n\nQue, para efectos de pago en el reconocimiento de los honorarios para el periodo académico corresponde de conformidad con el calendario académico emitido por el Consejo Académico mediante Resolución 112 de diciembre 18 de 2018 para las actividades del año 2019 que inicia en agosto  quince (15) y hasta diciembre dieciséis (16) del 2019, en los programas de posgrado en la Universidad, que para este caso se reconocerán de conformidad con los módulos de las horas lectivas efectivamente dictadas y aprobadas en los diferentes programas.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que realicen las actividades lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital.\n\nQue en materia de Seguridad y Salud en el Trabajo SST para docentes ocasionales de la Universidad Distrital se deben informar las obligaciones específicas se encuentra contemplado en el Decreto 1072 de 2015 artículo 2.2.4.2.2.16. y la resolución de rectoría No. 624 de 2017, el docente ocasional, hora catedra y honorarios debe cumplir con las normas del sistema general de riesgos laborales.\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de enero 31 de 2003, 0013-A de enero 31 de 2003, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. _______ del _______ de enero de 2019.\n\nQue, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "Vincular para el " + periodoStr + " Periodo Académico del año " + strconv.Itoa(vigencia) + " en los programas de posgrado de la Universidad Distrital Francisco José de Caldas, como docentes en la modalidad de Hora Cátedra de Vinculación Especial, en el escalafón y dedicación establecidas en la tabla, para el periodo comprendido entre el doce (12) de agosto al trece (13) de diciembre del 2019, a los siguientes docentes:"}
		}
		// accion = "Vincular"
		nombreDedicacion = "Hora Cátedra"
	case "HCH":
		if nivel == "PREGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\n Que la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), hora cátedra honorarios (HCH), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\n Que mediante Resolución Nro. 001 del 15 de febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\n Que los servicios de los docentes de vinculación especial, hora cátedra, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos, en los términos del artículo 73 de la Ley 30 de 1992.\n\n Que el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\n Que en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\n Que conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\n Que los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de enero 31 de 2003, 0013-A de enero 31 de 2003, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\n Que conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\n Que mediante Resolución 003 de enero 14 de 2019 y Resolución 021 de febrero 26 del 2019, por medio de las cuales expiden el calendario académico para el año 2019 para las actividades académicas tercer periodo académico del año 2019 que corresponde del doce (12) de agosto al trece (13) de diciembre del 2019.\n\n Que de conformidad con el artículo 2° del Decreto 1019 de seis (06) de junio de 2019 establece “A partir del 1° de enero de 2019, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en catorce mil doscientos diez pesos ($14.210) moneda corriente”.\n\n Que mediante resolución 312 del veinticinco (25) de julio del 2019 emitida por Rectoría de la Universidad Distrital Francisco José de Caldas en donde acoge y aplica, en lo pertinente única y expresamente al valor del punto para los docentes de Vinculación Especial hora cátedra, honorarios, Medio Tiempo Ocasional y Tiempo Completo Ocasional en pregrado en el valor de catorce mil doscientos diez pesos ($14.210) moneda corriente.\n\n Que en materia de Seguridad y Salud en el Trabajo SST para docentes ocasionales de la Universidad Distrital se deben informar las obligaciones específicas se encuentra contemplado en el Decreto 1072 de 2015 artículo 2.2.4.2.2.16. y la resolución de rectoría No. 624 de 2017, el docente ocasional, hora catedra y honorarios debe cumplir con las normas del sistema general de riesgos laborales.\n\n Que se hace necesario garantizar el reconocimiento de honorarios a docentes de hora cátedra en virtud de la ley 4 de 1992 que realicen las actividades de horas lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital.\n\n Que mediante Acuerdo 002 de marzo 17 de 2011 se modifica el artículo 43 del Acuerdo 011 de 2002, por el cual se pueden reconocer honorarios a personal pensionado en la modalidad de docentes de hora- cátedra por honorarios.\n\n Que para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. _______ del _______ de enero de 2019.\n\n Que, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "Reconocer Honorarios de conformidad con la ley 4 de 1992 para el " + periodoStr + " Periodo Académico del año " + strconv.Itoa(vigencia) + " en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, como docentes en la modalidad de Hora Cátedra Honorarios (HCH) de Vinculación Especial, en el escalafón y dedicación establecidas en la tabla, para el período comprendido entre el agosto doce (12) y hasta diciembre trece (13) del 2019, a los siguientes docentes:"}
		}
		if nivel == "POSGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que el estatuto general de la Universidad Distrital Francisco José de Caldas (Acuerdo 003 de Abril de 1997 del Consejo Superior Universitario), en el articulo 49, señala que los profesores de cátedra, ocasionales, visitantes y expertos, no son empleados públicos ni trabajadores de la Universidad, el reconociendo de sus servicios y prestaciones se hará mediante resolución.\n\nQue mediante Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las Universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital. (Artículos 3 y 4).\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes en las modalidades de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional, en virtud del árticulo 13 del Acuerdo 011 de Noviembre 15 de 2002 termino Fijo por periodos académicos.\n\nQue, el artículo 128 de la Carta Política consigna que nadie podrá desempeñar simultáneamente más de un empleo público ni recibir más de una asignación que provenga del tesoro público salvo las excepciones establecidas en la Ley.\n\nQue, en virtud de la anterior norma constitucional, el literal d del artículo 19 de la Ley 4 de 1992 determinó como excepción los honorarios percibidos por concepto hora cátedra.\n\nQue, mediante Acuerdo 006 de Julio 19 de 2002 se fija el valor de la Hora Cátedra por Honorarios y se establece un número máximo de horas para los docentes de carrera que presten servicios a la Universidad Distrital Francisco José de Caldas, en los programas de Postgrado, y modifican parcialmente los Acuerdos 005 y 007 de 2001.\n\nQue, mediante Acuerdo 005 de julio 27 de 2001, igualmente al Acuerdo 006 de Julio 19 de 2002 por medio del cual fija el valor Hora Cátedra Honorarios y se establece el número máximo de horas, para los docentes que presten servicios a la Universidad Distrital en los programas de Postgrado.\n\nQue, mediante Acuerdo 002 de Marzo 17 de 2011 se modifica el artículo 43 del Acuerdo 011 de 2002, por el cual se pueden vincular a personas pensionadas en la modalidad de docentes de hora-cátedra por honorarios.\n\nQue, para efectos de pago en el reconocimiento de los honorarios para el período académico corresponde de conformidad con el calendario académico emitido por el Consejo Académico mediante Resolución 112 de diciembre 18 de 2018 para las actividades del año 2019 que inicia en Agosto quince (15) y hasta Diciembre dieciséis (16) del 2019, en los programas de posgrado en la Universidad, que para este caso se reconocerán de conformidad con los módulos de las horas lectivas efectivamente dictadas y aprobadas en los diferentes programas.\n\nQue, mediante Resolución No. 001 del 15 de Febrero de 2012 emitida por la Vicerectoría Académica se establecen los procesos de selección y vinculación de docentes de vinculación especial.\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de Noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de Enero 31 de 2003, 0013-A de Enero 31 de 2003 Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\nQue, para efectos presupuestales de la presente Resolución se hará con cargo a la disponibilidad Presupuestales número _______ del _______ de ______________ del 2019.\n\nEn virtud de lo anteriormente expuesto,\n\n"}
			articulo = Articulo{Texto: "Reconocer Honorarios de conformidad con la ley 4 de 1992 para el " + periodoStr + " Periodo Académico del año " + strconv.Itoa(vigencia) + " en los programas de posgrado de la Universidad Distrital Francisco José de Caldas, como docentes en la modalidad de Hora Cátedra Honorarios (HCH) de Vinculación Especial, en el escalafón y dedicación establecidas en la tabla, para el período comprendido entre el agosto quince (15) y hasta diciembre dieciséis (16) del 2019, a los siguientes docentes:"}
		}
		// accion = "Reconocer Honorarios"
		nombreDedicacion = "Hora Cátedra Honorarios"
	case "TCO-MTO":
		resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), Honorarios hora catedra, en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\nQue mediante Resolución Nro. 001 del 15 de febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial, hora cátedra, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos, en los términos del artículo 73 de la Ley 30 de 1992.\n\nQue el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue mediante Resolución 003 de enero 14 de 2019 y Resolución 021 de febrero 26 del 2019, por medio de las cuales expiden el calendario académico para el año 2019 para las actividades académicas tercer periodo académico del año 2019 que corresponde del doce (12) de agosto al trece (13) de diciembre del 2019.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que realicen las actividades lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital.\n\nQue de conformidad con el artículo 2° del Decreto 1019 de seis (06) de junio de 2019 establece “A partir del 1° de enero de 2019, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en catorce mil doscientos diez pesos ($14.210) moneda corriente”.\n\nQue mediante resolución 312 del veinticinco (25) de julio del 2019 emitida por Rectoría de la Universidad Distrital Francisco José de Caldas en donde acoge y aplica, en lo pertinente única y expresamente al valor del punto para los docentes de Vinculación Especial hora cátedra, honorarios, Medio Tiempo Ocasional y Tiempo Completo Ocasional en pregrado en el valor de catorce mil doscientos diez pesos ($14.210) moneda corriente.\n\nQue en materia de Seguridad y Salud en el Trabajo SST para docentes ocasionales de la Universidad Distrital se deben informar las obligaciones específicas se encuentra contemplado en el Decreto 1072 de 2015 artículo 2.2.4.2.2.16. y la resolución de rectoría No. 624 de 2017, el docente ocasional, hora catedra y honorarios debe cumplir con las normas del sistema general de riesgos laborales.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. _______ del __________ de enero de 2019.\n\nQue, en mérito de lo expuesto,\n\n"}
		articulo = Articulo{Texto: "Vincular para el tercer Periodo académico del año 2019 en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, como docentes en la modalidad de Hora Cátedra de Vinculación Especial, en el escalafón y dedicación establecidas en la tabla, para el periodo comprendido entre el doce (12) de agosto al trece (13) de diciembre del 2019, a los siguientes docentes:"}
		// accion = "Vincular"
		nombreDedicacion = "Medio Tiempo Ocasional y Tiempo Completo Ocasional"
	}

	if tipo != "1" {
		articulo = Articulo{Texto: "Modificar la Resolución No " + numero + " del " + strconv.Itoa(vigencia) + " en cuanto al número de horas semanales y el valor total para el " + periodoStr + " Período Académico del " + strconv.Itoa(vigencia) + ", como docentes en la modalidad de " + nombreDedicacion + " de Vinculación Especial, en el escalafón y dedicación establecidas en la siguiente tabla:"}
	}

	articulos = append(articulos, articulo)

	/////ARTICULO 2
	if (dedicacion == "HCH" && nivel == "POSGRADO") {
		articulo = Articulo{Texto: "El pago de los honorarios por los servicios prestados a los profesores de Hora Cátedra por Honorarios según su escalafón, se cancelará previa certificación de las horas efectivamente dictadas, expedida por el Decano (a)."}

	}else{
		if(dedicacion == "HCP" && nivel == "POSGRADO"){
			articulo = Articulo{Texto: "El pago de los servicios prestados por los profesores de vinculación especial a que se refiere el artículo anterior, según su escalafón, se hará previa certificación de las horas efectivamente dictadas, que se encuentren inmersas en el correspondiente plan de trabajo de la gestión académica, expedida por el decano y/o director de proyecto curricular.\n\n"}
		}else{
			articulo = Articulo{Texto: "El pago de los servicios prestados por los profesores de vinculación especial a que se refiere el artículo anterior, según su escalafón, se hará  previa certificación de las horas efectivamente dictadas, que se encuentren inmersas en el correspondiente plan de trabajo de la gestión académica, expedida por el decano y/o director de proyecto curricular.\n\n"}
		    paragrafo := Paragrafo{Texto: "El valor del punto en pesos para el reconocimiento y pago de los docentes de hora cátedra, será el que fije el Gobierno Nacional mediante decreto, cada año, y que la Universidad Distrital Francisco José de Caldas acoja mediante acto administrativo, respecto de los docentes de vinculación especial."}
            articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		}
	}

	articulos = append(articulos, articulo)
  /////ARTICULO 3
  if (dedicacion == "HCH" && nivel == "POSGRADO") {
		articulo = Articulo{Texto: "Los docentes deberán cumplir con las obligaciones inherentes a la naturaleza del servicio, contempladas en la Ley, en los Reglamentos de la Universidad y en los Planes de Trabajo Entregados por el profesor y aprobados por Decano (a).\n\n"}
		paragrafo = Paragrafo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad mediante acto administrativo hará la liquidación con corte a la fecha del cumplido expedido por el Decano (a).\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "El docente cumplirá con lo establecido en la circular No. 004 de Julio 19 de 2017 de Rectoría con relación al pago de aportes a seguridada social de forma independiente de conformidad con la ley."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}else{
		articulo = Articulo{Texto: "Los docentes en cuestión deberán cumplir con las obligaciones inherentes a la naturaleza del servicio, contempladas en la ley, en los reglamentos de la Universidad Distrital y en los planes de trabajo establecidos en el aplicativo de gestión académica, entregados por cada profesor, y aprobados por el coordinador del correspondiente programa académico, decano y/o director.\n\n"}
		if dedicacion == "TCO-MTO" {
			paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el Plan de Trabajo del período académico correspondiente."}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		} else {
			if dedicacion == "HCP" {
				paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el Plan de Trabajo del período académico correspondiente."}
			    articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
			}else{
				paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el Plan de Trabajo del período académico correspondiente.\n\n"}
				articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
				paragrafo = Paragrafo{Texto: "El docente cumplirá con lo establecido en la circular No. 004 de Julio 19 de 2017 de Rectoría con relación al pago de aportes a seguridada social de forma independiente de conformidad con la ley."}
			    articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
			}
		}
	}
	articulos = append(articulos, articulo)
  ////ARTICULO 4
	if (dedicacion == "HCH" && nivel == "POSGRADO") {
		articulo = Articulo{Texto: "El gasto que ocasione la presente resolución se hará cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal.\n\n"}
		paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias realizadas por la Secretaría de Hacienda Distrital y se realizarán dentro de los primeros cinco días de cada mes."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		articulos = append(articulos, articulo)
	}else{
		if dedicacion == "HCP" {
			articulo = Articulo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos."}
			articulos = append(articulos, articulo)
		}
		if dedicacion == "HCH" {
			articulo = Articulo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del reconocimiento con corte a la fecha del respectivo cumplido, cancelando los correspondientes Honorarios, conforme al cálculo que efectúe la División de Recursos Humanos."}
			articulos = append(articulos, articulo)
		}
		if dedicacion == "TCO-MTO" {
			articulo = Articulo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos."}
			articulos = append(articulos, articulo)
		}
	}
 /////ARTICULO 5
  if (dedicacion == "HCH" && nivel == "POSGRADO") {
			articulo = Articulo{Texto: "Bajo el supuesto que, se declare la suspensión de actividades académicas por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial la obligación de prestar sus servicios y para la Universidad, la de pagar los salarios y prestaciones correspondientes al periodo suspendido, pero persistirá para esta última, la de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda.\n\nSi la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas mediante acto administrativo motivado, podrá declarar la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, conforme al cálculo que efectúe la División de Recursos Humanos."}
	}else{
		  articulo = Articulo{Texto: "El gasto que ocasione el presente acto administrativo, se hará con cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal.\n\n"}
		  if(dedicacion == "HCH" && nivel == "PREGRADO"){
		  	paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias de la Secretaría de Hacienda Distrital y se realizarán dentro de los primeros cinco días de cada mes."}
		  	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		  }else{
		  	paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias realizadas por la Secretaría de Hacienda Distrital y se realizarán dentro de los primeros cinco días de cada mes."}
		  	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		  }
	}
	articulos = append(articulos, articulo)
	/////ARTICULO 6
	if (dedicacion == "HCH" && nivel == "POSGRADO") {
		articulo = Articulo{Texto: "El Decreto 1072 de 2015 artículo 2.2.4.2.2.16. y la resolución de rectoría No. 624 de 2017, el docente ocasional, hora catedra y honorarios debe dar cumplimiento con las normas del sistema general de riesgos laborales las cuales deberá consultar en la mencionada norma."}
		articulos = append(articulos, articulo)
	}else{
		if dedicacion == "HCH" {
			articulo = Articulo{Texto: "Bajo el supuesto que, se declare la suspensión de actividades académicas por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial la obligación de prestar sus servicios y para la Universidad, la de reconocer los honorarios correspondientes al periodo suspendido.\n\nSi la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas mediante acto administrativo motivado, podrá declarar la terminación con corte a la fecha del respectivo cumplido, cancelando honorarios correspondientes, conforme al cálculo que efectúe la División de Recursos Humanos."}
			articulos = append(articulos, articulo)
		} else {
			articulo = Articulo{Texto: "Bajo el supuesto que, se declare la suspensión de actividades académicas por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial la obligación de prestar sus servicios y para la Universidad, la de pagar los salarios y prestaciones correspondientes al periodo suspendido, pero persistirá para esta última, la de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda.\n\nSi la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas mediante acto administrativo motivado, podrá declarar la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos."}
			articulos = append(articulos, articulo)
		}
	}
  /////ARTICULO 7
  	if (dedicacion == "HCH" && nivel == "PREGRADO") {
		articulo = Articulo{Texto: "El Decreto 1072 de 2015 artículo 2.2.4.2.2.16. y la resolución de rectoría No. 624 de 2017, el docente ocasional, hora catedra y honorarios debe dar cumplimiento con las normas del sistema general de riesgos laborales las cuales deberá consultar en la mencionada norma."}
		articulos = append(articulos, articulo)
	}else{
		if dedicacion == "TCO-MTO" {
			articulo = Articulo{Texto: "El Decreto 1072 de 2015 artículo 2.2.4.2.2.16. y la resolución de rectoría No. 624 de 2017, el docente ocasional, hora catedra y honorarios debe dar cumplimiento con las normas del sistema general de riesgos laborales las cuales deberá consultar en la mencionada norma."}
		    articulos = append(articulos, articulo)
		}else{
			if dedicacion == "HCP" {
				articulo = Articulo{Texto: "El Decreto 1072 de 2015 artículo 2.2.4.2.2.16. y la resolución de rectoría No. 624 de 2017, el docente ocasional, hora catedra y honorarios debe dar cumplimiento con las normas del sistema general de riesgos laborales las cuales deberá consultar en la mencionada norma."}
		        articulos = append(articulos, articulo)
			}else{
				articulo = Articulo{Texto: "Comuníquese la presente resolución a los docentes mencionados en el artículo 1º, quienes deberán manifestar, bajo la gravedad de juramento que se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, que no tienen cruces de horarios ni ostentan otra vinculación de carácter público, diferente a hora cátedra, en entidades de educación oficiales, siempre y cuando los honorarios no sumen más de ocho (8) horas diarias de trabajo a varias entidades."}
	    	    articulos = append(articulos, articulo)
			}
	    }
	}
	/////ARTICULO 8
	if (dedicacion == "HCH" && nivel == "POSGRADO") {
		articulo = Articulo{Texto: "El presente acto administrativo se expide a los ______________(__) días del mes de __________ del año " + strconv.Itoa(vigencia) + " y surte efectos de conformidad con la resolución 112 de diciembre 18 de 2018 por medio del cual se expide el calendario académico para el año 2019 para las actividades académicas  " + periodoStr + " Periodo Académico para el año " + strconv.Itoa(vigencia) + " que corresponde del agosto quince (15) y hasta diciembre dieciséis (16) del 2019."}
		articulos = append(articulos, articulo)
	}else{
		if (dedicacion == "HCH" && nivel == "PREGRADO") {
			articulo = Articulo{Texto: "Comuníquese la presente resolución a los docentes mencionados en el artículo 1º, quienes deberán manifestar, bajo la gravedad de juramento que se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, que no tienen cruces de horarios ni ostentan otra vinculación de carácter público, diferente a hora cátedra, en entidades de educación oficiales, siempre y cuando los honorarios no sumen más de ocho (8) horas diarias de trabajo a varias entidades."}
	        articulos = append(articulos, articulo)
		}else{
			if dedicacion == "TCO-MTO" {
				articulo = Articulo{Texto: "Comuníquese la presente resolución a los docentes mencionados en el artículo 1º, quienes deberán manifestar, bajo la gravedad de juramento que se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, que no tienen cruces de horarios ni ostentan otra vinculación de carácter público, diferente a hora cátedra, en entidades de educación oficiales, siempre y cuando los honorarios no sumen más de ocho (8) horas diarias de trabajo a varias entidades."}
	            articulos = append(articulos, articulo)
			}else{
				articulo = Articulo{Texto: "Comuníquese la presente resolución a los docentes mencionados en el artículo 1º, quienes deberán manifestar, bajo la gravedad de juramento que se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, que no tienen cruces de horarios ni ostentan otra vinculación de carácter público, diferente a hora cátedra, en entidades de educación oficiales, siempre y cuando los honorarios no sumen más de ocho (8) horas diarias de trabajo a varias entidades."}
		        articulos = append(articulos, articulo)
			}
		}
	}
	/////ARTICULO 9
	if (dedicacion == "HCH" && nivel == "PREGRADO") {
		articulo = Articulo{Texto: "El presente acto administrativo se expide a los ______________(__) días del mes de __________ del año " + strconv.Itoa(vigencia) + " y surte efectos de conformidad con la resolución 003 de enero catorce (14) de 2019, por medio del cual se expide el calendario académico para el año 2019 para las actividades académicas  " + periodoStr + " Periodo Académico para el año " + strconv.Itoa(vigencia) + " que corresponde del doce (12) de agosto al trece (13) diciembre del 2019."}
		articulos = append(articulos, articulo)
	}
	if dedicacion == "TCO-MTO" {
		articulo = Articulo{Texto: "El presente acto administrativo se expide a los ______________(__) días del mes de __________ del año " + strconv.Itoa(vigencia) + " y surte efectos de conformidad con la resolución 003 de enero catorce (14) de 2019, por medio del cual se expide el calendario académico para el año 2019 para las actividades académicas  " + periodoStr + " Periodo Académico para el año " + strconv.Itoa(vigencia) + " que corresponde del doce (12) de agosto al trece (13) diciembre del 2019."}
		articulos = append(articulos, articulo)
	}
	if (dedicacion == "HCP" && nivel == "PREGRADO") {
		articulo = Articulo{Texto: "El presente acto administrativo se expide a los ______________(__) días del mes de __________ del año " + strconv.Itoa(vigencia) + " y surte efectos de conformidad con la resolución 003 de enero catorce (14) de 2019, por medio del cual se expide el calendario académico para el año 2019 para las actividades académicas  " + periodoStr + " Periodo Académico para el año " + strconv.Itoa(vigencia) + " que corresponde del doce (12) de agosto al trece (13) diciembre del 2019."}
		articulos = append(articulos, articulo)
	}
	if (dedicacion == "HCP" && nivel == "POSGRADO") {
		articulo = Articulo{Texto: "El presente acto administrativo se expide a los ______________(__) días del mes de __________ del año " + strconv.Itoa(vigencia) + " y surte efectos de conformidad con la resolución 112 de diciembre 18 de 2018, por medio del cual se expide el calendario académico para el año 2019 para las actividades académicas  " + periodoStr + " Periodo Académico para el año " + strconv.Itoa(vigencia) + " que corresponde del agosto quince (15) y hasta diciembre dieciséis (16) del 2019."}
		articulos = append(articulos, articulo)
	}

	resolucion.Articulos = articulos //articulos//articulos

	return resolucion
}
