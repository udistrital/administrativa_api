package models

import (
	"fmt"
	"strconv"

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
	Vinculacion             ResolucionVinculacionDocente
	Consideracion           string
	Preambulo               string
	Vigencia                int
	Numero                  string
	Id                      int
	Articulos               []Articulo
	Titulo                  string
	CuadroResponsabilidades string
}

func GetOneResolucionCompleta(idResolucion string) (resolucion ResolucionCompleta) {
	o := orm.NewOrm()
	var temp []Resolucion
	idRes, _ := strconv.Atoi(idResolucion)

	_, err := o.QueryTable("resolucion").Filter("id_resolucion", idRes).All(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	resolucionCompleta := ResolucionCompleta{Id: temp[0].Id, Consideracion: temp[0].ConsideracionResolucion, Preambulo: temp[0].PreambuloResolucion, Vigencia: temp[0].Vigencia, Numero: temp[0].NumeroResolucion, Titulo: temp[0].Titulo, CuadroResponsabilidades: temp[0].CuadroResponsabilidades}

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
		v.CuadroResponsabilidades = m.CuadroResponsabilidades
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

func GetTemplateResolucion(dedicacion, nivel, periodo, tipo string) (res ResolucionCompleta) {
	var resolucion ResolucionCompleta
	var articulos []Articulo
	var articulo Articulo
	var paragrafo Paragrafo
	//var vigencia, _, _ = time.Now().Date()
	//var accion string
	//var periodoStr string
	//var nombreDedicacion string
	/*
		switch periodo {
		case "1":
			periodoStr = "primer"
		case "2":
			periodoStr = "segundo"
		case "3":
			periodoStr = "tercer"
		}
	*/

	switch dedicacion {
	case "HCP":
		if nivel == "PREGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que el Decreto 1279 de 2002, mediante el cual se establece el régimen salarial y prestacional de los docentes de las universidades estatales, señala en el artículo 3º que: “(…) los profesores ocasionales no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto”, precisando que, “no obstante, su vinculación se hace conforme a las reglas que define cada Universidad, con sujeción a lo dispuesto por la Ley 30 de 1992 y demás disposiciones constitucionales y legales vigentes”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de Tiempo Completo Ocasional (TCO), Medio Tiempo Ocasional (MTO), Hora Cátedra (HC) y Hora Catedra Honorarios (HCH), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de 2002 del Consejo Superior Universitario (Estatuto Docente), a término fijo por períodos académicos.\n\nQue mediante Resolución 001 de 15 de febrero de 2012 de la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la institución de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial deberán ser reconocidos en los términos del inciso 2º del artículo 74 de la Ley 30 de 1992, esto es, “mediante resolución”.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de 18 de diciembre de 2018 del Departamento Administrativo de la Función Pública, “en el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración (…)”.\n\nQue conforme al parágrafo 1º del artículo 5º de la Resolución 001 de 2012 de la Vicerrectoría Académica, establece que, “(…) Para efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue de conformidad con el artículo 2° del Decreto Nacional 447 de 29 de marzo 2022, “(…) a partir del 1° de enero de 2022 se fija el valor del punto salarial para los empleados públicos docentes a quienes se les aplica el Decreto 1279 de 2002 y demás disposiciones que lo modifiquen o adicionen en dieciséis mil cuatrocientos cuarenta y un pesos ($16.441) moneda corriente (…)”.\n\nQue en virtud del mencionado decreto, se acoge y aplica, en lo pertinente, única y expresamente al valor del punto salarial en DIECISÉIS MIL CUATROCIENTOS CUARENTA Y UN PESOS ($16.441) MONEDA CORRIENTE, para los docentes de Vinculación Especial Hora Cátedra para los programas de pregrado, esto en concordancia con el artículo 2° del Acuerdo 012 de 2002 del Consejo Superior Universitario.\n\nQue la Resolución de Rectoría 087 de 2022, “Por medio de la cual se adoptan las medidas para el regreso del personal administrativo de la Universidad Distrital Francisco José de Caldas a las actividades laborales presenciales”, determina: “El reinicio total de actividades administrativas y académico- administrativas presenciales será a partir del 02 de marzo del 2022. La totalidad de servidores administrativos (libre nombramiento y remoción, carrera administrativa, provisionalidad y temporalidad) deberán asistir diariamente a desempeñar sus funciones de manera presencial en las correspondientes sedes de la Universidad Distrital Francisco José de Caldas”.\n\nQue igualmente, el artículo 4 del mentado acto administrativo señala que “(…) El regreso a las actividades presenciales implica el cumplimiento cabal de la Resolución 777 del 2021 del Ministerio de Salud y Protección Social, así como del Protocolo de Bioseguridad de la Universidad-2022”.\n\nQue en materia de Seguridad y Salud en el Trabajo (SST) para docentes ocasionales de la Universidad Distrital Francisco José de Caldas, se deben informar las obligaciones específicas de que trata el artículo 2.2.4.2.2.16. del Decreto Nacional 1072 de 2015, así como la Resolución de Rectoría 624 de 2017, conforme a los cuales, el docente ocasional, de hora catedra y por honorarios, debe cumplir con las normas del Sistema General de Riesgos Laborales.\n\nQue el Consejo Académico, en uso de las funciones estatutarias establecidas en los literales b) y e) del artículo 18 del Estatuto General -Acuerdo 003 de 1997 del CSU-, mediante Resolución 13 de 15 de marzo de 2022, fijó el calendario académico especial para los periodos académicos 2022-I y 2022-III en los distintos programas de pregrado de la Universidad Distrital Francisco José de Caldas.\n\nQue a través de la Resolución 014 de 29 de marzo de 2022, el Consejo Académico modificó la Resolución 013 de 15 de marzo de 2022, por medio de la cual se fijó el calendario académico especial para los periodos académicos 2022-1 y 2022-lII en los distintos programas académicos de pregrado de la Universidad Distrital Francisco José de Caldas\".\n\nQue descrito lo anterior, las actividades académicas correspondientes al período académico 2022-I se desarrollarán entre el 26 de abril y el 19 de agosto de 2022; y las actividades académicas correspondientes al periodo académico 2022-III se desarrollarán entre el 30 de agosto y el 24 de diciembre de 2022.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que realicen las actividades lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal (Sueldo Básico, Prima Vacaciones, Prima Navidad, Prima de Servicios y Cesantías) 0xxx del xx de enero de 2022.\n\nQue, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "VINCULACIÓN. Vincular a los siguientes docentes para los periodos académicos 2022-I y 2022-III, comprendidos entre el 26 de abril y el 19 de agosto de 2022, y entre el 30 de agosto de 2022 y el 24 de diciembre de 2022, respectivamente, en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, en la modalidad de Hora Cátedra Prestaciones, en el escalafón y dedicación establecidas en la siguiente tabla, y con cargo a los recursos asignados en el Certificado de Disponibilidad Presupuestal (Sueldo Básico, Prima Vacaciones, Prima Navidad, Prima de Servicios y Cesantías) xxx del xx de enero de 2022:\n\n"}
			paragrafo := Paragrafo{Texto: "La vinculación de los docentes referidos en este artículo se suspenderá durante el periodo comprendido entre el 20 y el 29 de agosto de 2022, el cual corresponde a la terminación del periodo académico 2022-I y el inicio del periodo académico 2022-III. Por lo tanto, en ese lapso no existirá para los docentes la obligación de prestar sus servicios, ni para la Universidad la obligación de pagar salarios y prestaciones; pero persistirá la obligación de la institución de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda."}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
			paragrafo = Paragrafo{Texto: "En todo caso, la continuidad de la vinculación de los docentes mencionados en este artículo para el periodo académico 2022-III estará sujeta a las variaciones que puedan sufrir las cargas académicas, de acuerdo con las dinámicas institucionales relativas a la matrícula y la conformación de los grupos, entre otras. En caso de que las necesidades del servicio para el periodo académico 2022-III no demanden la vinculación del mismo número de docentes que para el periodo académico 2022-I, la Universidad podrá dar por terminada la vinculación de aquellos docentes cuyos servicios no sean requeridos, y efectuar la liquidación correspondiente con corte a la fecha del respectivo cumplido, conforme al cálculo que efectúe la División de Recursos Humanos."}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		}
		if nivel == "POSGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que el artículo 3º del Decreto 1279 de 2002, mediante el cual se establece el régimen salarial y prestacional de los docentes de las universidades, señala que “(…)los profesores ocasionales no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto”, precisando que, “no obstante, su vinculación se hace conforme a las reglas que define cada Universidad, con sujeción a lo dispuesto por la ley 30 de 1992 y demás disposiciones constitucionales y legales vigentes”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de: Tiempo Completo Ocasional (TCO), Medio Tiempo Ocasional (MTO), Hora Cátedra (HC) y Hora Catedra Honorarios (HCH), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de 2002 del Consejo Superior Universitario (Estatuto Docente), a término fijo o por períodos académicos.\n\nQue mediante Resolución 001 de 15 de febrero de 2012 de la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la institución de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial deberán ser reconocidos en los términos del inciso 2º del artículo 74 de la Ley 30 de 1992, esto es, “mediante resolución”.\n\nQue la Corte Constitucional en la Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de 18 de diciembre de 2018 del Departamento Administrativo de la Función Pública, “en el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo 5° de la Resolución 001 de 2012 de la Vicerrectoría Académica, “(p)ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue el Consejo Académico, en ejercicio de las funciones establecidas en los literales b) y e) del artículo 18 del Estatuto General de la Universidad Distrital Francisco José de Caldas (Acuerdo 003 de abril 08 de 1997), , aprobó el Calendario Académico del año 2022, para los programas académicos de posgrado de la Universidad Distrital Francisco José de Caldas, mediante Resolución 051 de 28 de septiembre de 2021.\n\nQue el artículo1° de la Resolución 051 de 28 de septiembre de 2021 determinó que el inicio de las clases en los programas de posgrado puede tener lugar en el periodo comprendido entre el 20 de enero y hasta el 19 de febrero de 2022, habiéndose establecido el cierre del periodo para el periodo comprendido entre el 4 de junio y el 24 de junio de 2022.\n\nQue el artículo 1° de la Resolución de Rectoría 210 de 2021, determina: “Autorizar a las directivas académico – administrativas de la Universidad Distrital Francisco José de Caldas, a dar inicio al plan de retorno seguro y gradual en alternancia de las dependencias administrativas, académico administrativas y comunidad académica, de conformidad con los criterios, condiciones y protocolos de bioseguridad establecidas por las autoridades sanitarias, en el marco de la fase de Aislamiento Selectivo, Distanciamiento Individual Responsable y Reactivación Económica Segura de la pandemia por COVID-19”.\n\nQue se hace necesario garantizar la vinculación de docentes Hora Cátedra Prestaciones para que realicen las actividades lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital Francisco José de Caldas.\n\nQue, en materia de Seguridad y Salud en el Trabajo (SST), para docentes ocasionales de la Universidad Distrital Francisco José de Caldas, se deben informar las obligaciones específicas contempladas en el Decreto Nacional 1072 de 2015 (art. 2.2.4.2.2.16.) y en la Resolución de Rectoría 624 de 2017.\n\nQue, para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. XXXX del XX de enero de 2022.\n\nQue en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "VINCULACIÓN. Vincular a los siguientes docentes en la modalidad de Hora Cátedra Prestaciones, conforme al escalafón y dedicación establecidas en la siguiente tabla, en los programas de posgrado de la Universidad Distrital Francisco José de Caldas para el periodo académico de 2022-1, comprendido entre el xxx de xxxxxx y hasta el xxx de xxxxxxxx de 2022, con cargo a los recursos asignados en el Certificado de Disponibilidad Presupuestal xxx del xx de enero de 2022 y de conformidad con la Ley 4ª de 1992:\n\n"}
			paragrafo := Paragrafo{Texto: "El pago de los servicios prestados por los docentes anteriormente relacionados, se hará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal (Sueldo Básico, Prima Vacaciones, Prima Navidad, Prima de Servicios y Cesantías) XXX del XX de enero de 2022.\n\n"}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		}
		// accion = "Vincular"
		//nombreDedicacion = "Hora Cátedra"
	case "HCH":
		if nivel == "PREGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que el Decreto 1279 de 2002, mediante el cual se establece el régimen salarial y prestacional de los docentes de las universidades estatales, señala en el artículo 3º que: “(…) los profesores ocasionales no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto”, precisando que, “no obstante, su vinculación se hace conforme a las reglas que define cada Universidad, con sujeción a lo dispuesto por la Ley 30 de 1992 y demás disposiciones constitucionales y legales vigentes”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de Tiempo Completo Ocasional (TCO), Medio Tiempo Ocasional (MTO), Hora Cátedra (HC) y Hora Catedra Honorarios (HCH), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de 2002 del Consejo Superior Universitario (Estatuto Docente), a término fijo por períodos académicos.\n\nQue mediante Resolución 001 de 15 de febrero de 2012 de la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la institución de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial deberán ser reconocidos en los términos del inciso 2º del artículo 74 de la Ley 30 de 1992, esto es, “mediante resolución”.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de 18 de diciembre de 2018 del Departamento Administrativo de la Función Pública, “en el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración (…)”.\n\nQue conforme al parágrafo 1º del artículo 5º de la Resolución 001 de 2012 de la Vicerrectoría Académica, establece que, “(…) Para efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue de conformidad con el artículo 2° del Decreto Nacional 447 de 29 de marzo 2022, “(…) a partir del 1° de enero de 2022 se fija el valor del punto salarial para los empleados públicos docentes a quienes se les aplica el Decreto 1279 de 2002 y demás disposiciones que lo modifiquen o adicionen en dieciséis mil cuatrocientos cuarenta y un pesos ($16.441) moneda corriente (…)”.\n\nQue en virtud del mencionado decreto, se acoge y aplica, en lo pertinente, única y expresamente al valor del punto salarial en DIECISÉIS MIL CUATROCIENTOS CUARENTA Y UN PESOS ($16.441) MONEDA CORRIENTE, para los docentes de Vinculación Especial Hora Cátedra para los programas de pregrado, esto en concordancia con el artículo 2° del Acuerdo 012 de 2002 del Consejo Superior Universitario.\n\nQue la Resolución de Rectoría 087 de 2022, “Por medio de la cual se adoptan las medidas para el regreso del personal administrativo de la Universidad Distrital Francisco José de Caldas a las actividades laborales presenciales”, determina: “El reinicio total de actividades administrativas y académico- administrativas presenciales será a partir del 02 de marzo del 2022. La totalidad de servidores administrativos (libre nombramiento y remoción, carrera administrativa, provisionalidad y temporalidad) deberán asistir diariamente a desempeñar sus funciones de manera presencial en las correspondientes sedes de la Universidad Distrital Francisco José de Caldas”.\n\nQue igualmente, el artículo 4 del mentado acto administrativo señala que “(…) El regreso a las actividades presenciales implica el cumplimiento cabal de la Resolución 777 del 2021 del Ministerio de Salud y Protección Social, así como del Protocolo de Bioseguridad de la Universidad-2022”.\n\nQue en materia de Seguridad y Salud en el Trabajo (SST) para docentes ocasionales de la Universidad Distrital Francisco José de Caldas, se deben informar las obligaciones específicas de que trata el artículo 2.2.4.2.2.16. del Decreto Nacional 1072 de 2015, así como la Resolución de Rectoría 624 de 2017, conforme a los cuales, el docente ocasional, de hora catedra y por honorarios, debe cumplir con las normas del Sistema General de Riesgos Laborales.\n\nQue el Consejo Académico, en uso de las funciones estatutarias establecidas en los literales b) y e) del artículo 18 del Estatuto General -Acuerdo 003 de 1997 del CSU-, mediante Resolución 13 de 15 de marzo de 2022, fijó el calendario académico especial para los periodos académicos 2022-I y 2022-III en los distintos programas de pregrado de la Universidad Distrital Francisco José de Caldas.\n\nQue a través de la Resolución 014 de 29 de marzo de 2022, el Consejo Académico modificó la Resolución 013 de 15 de marzo de 2022, por medio de la cual se fijó el calendario académico especial para los periodos académicos 2022-1 y 2022-lII en los distintos programas académicos de pregrado de la Universidad Distrital Francisco José de Caldas\".\n\nQue descrito lo anterior, las actividades académicas correspondientes al período académico 2022-I se desarrollarán entre el 26 de abril y el 19 de agosto de 2022; y las actividades académicas correspondientes al periodo académico 2022-III se desarrollarán entre el 30 de agosto y el 24 de diciembre de 2022.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que realicen las actividades lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal (Sueldo Básico, Prima Vacaciones, Prima Navidad, Prima de Servicios y Cesantías) 0xxx del xx de enero de 2022.\n\nQue, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "VINCULACIÓN. Vincular a los siguientes docentes para los periodos académicos 2022-I y 2022-III, comprendidos entre el 26 de abril y el 19 de agosto de 2022, y entre el 30 de agosto de 2022 y el 24 de diciembre de 2022, respectivamente, en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, en la modalidad de Hora Cátedra Prestaciones, en el escalafón y dedicación establecidas en la siguiente tabla, y con cargo a los recursos asignados en el Certificado de Disponibilidad Presupuestal (Sueldo Básico, Prima Vacaciones, Prima Navidad, Prima de Servicios y Cesantías) xxx del xx de enero de 2022:\n\n"}
			paragrafo := Paragrafo{Texto: "La vinculación de los docentes referidos en este artículo se suspenderá durante el periodo comprendido entre el 20 y el 29 de agosto de 2022, el cual corresponde a la terminación del periodo académico 2022-I y el inicio del periodo académico 2022-III. Por lo tanto, en ese lapso no existirá para los docentes la obligación de prestar sus servicios, ni para la Universidad la obligación de pagar salarios y prestaciones; pero persistirá la obligación de la institución de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda."}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
			paragrafo = Paragrafo{Texto: "En todo caso, la continuidad de la vinculación de los docentes mencionados en este artículo para el periodo académico 2022-III estará sujeta a las variaciones que puedan sufrir las cargas académicas, de acuerdo con las dinámicas institucionales relativas a la matrícula y la conformación de los grupos, entre otras. En caso de que las necesidades del servicio para el periodo académico 2022-III no demanden la vinculación del mismo número de docentes que para el periodo académico 2022-I, la Universidad podrá dar por terminada la vinculación de aquellos docentes cuyos servicios no sean requeridos, y efectuar la liquidación correspondiente con corte a la fecha del respectivo cumplido, conforme al cálculo que efectúe la División de Recursos Humanos."}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		}
		if nivel == "POSGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional, de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo, en su artículo 3º, que “(…) los profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\nQue el artículo 128 de la Carta Política establece que “nadie podrá desempeñar simultáneamente más de un empleo público ni recibir más de una asignación que provenga del tesoro público salvo las excepciones establecidas en la Ley (…)”.\n\nQue en consonancia con lo anterior, la Ley 4ª de 1992, en su artículo 19, determinó que “(…) nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del Tesoro Público, o de empresas o de instituciones en las que tenga parte mayoritaria el Estado. Exceptúense las siguientes asignaciones: (…) d. Los honorarios percibidos por concepto de hora-cátedra”.\n\nQue junto a lo anterior, el parágrafo de la norma en cita establece que “(…) no se podrán recibir honorarios que sumados correspondan a más de ocho (8) horas diarias de trabajo a varias entidades”.\n\nQue sobre el mismo tema, la Corte Constitucional, en sentencia C-133 de 1993, señaló que, “(…) si bien es cierto que en el artículo 128 C.P. se consagra una incompatibilidad, no lo es menos que ésta se encuentra en íntima relación de conexidad con la remuneración de los servidores estatales; basta ver que en ella se prohíbe la concurrencia de dos o más cargos públicos en una misma persona, tanto como recibir más de una asignación que provenga del erario público. El término ‘asignación´ comprende toda clase de remuneración que emane del tesoro público, llámese sueldo, honorario, mesada pensional, etc. (…)”.\n\nQue de otra parte, conforme a las normas y la jurisprudencia expuestas, la persona que tiene la calidad de pensionado del sector público podrá percibir otra asignación del Tesoro Público, siempre que la misma provenga de las excepciones establecidas en el artículo 19 de la Ley 4ª de 1992, como es el caso de los honorarios percibidos por los docentes que presten el servicio a una universidad o institución educativa estatal, mediante el sistema de “hora cátedra honorarios” (HCH).\n\nQue mediante Resolución 001 del 15 de febrero de 2012 dela Vicerrectoría Académica, se estableció el procedimiento para la selección y vinculación a la Universidad de docentes de vinculación especial.\n\nQue se hace necesario garantizar el reconocimiento de honorarios a docentes de hora cátedra en virtud de la Ley 4ª de 1992, que realicen actividades de horas lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital Francisco José de Caldas.\n\nQue, mediante el artículo 2° del Acuerdo 006 del 2002, que modificó el artículo 1° del Acuerdo 007 del 2001, ambos del Consejo Superior Universitario, se estableció que a los docentes hora catedra que pertenecen a la carrera docente de la Universidad Distrital Francisco José de Caldas se les podrá reconocer hasta un máximo de seis (6) horas semanales adicionales a su carga normal, en posgrado.\n\nQue, mediante Acuerdo 05 de 2001, el Consejo Superior Universitario fijó el valor de la hora cátedra y estableció el número máximo de horas para los docentes que prestan servicios a la Universidad Distrital Francisco José de Caldas en los programas de posgrado, estableciendo que su liquidación se efectúa con base en el salario mínimo mensual legal vigente.\n\nQue los docentes de carrera de la entidad, en todo caso, deberán cumplir, en el correspondiente plan de trabajo, con la carga académica de doce (12) horas semanales, previa aprobación, por parte, de los correspondientes decano y coordinador.\n\nQue, a los funcionarios estatales y pensionados del sector público, se les vinculará a la institución como docentes en la modalidad “hora cátedra honorarios” (HCH), los cuales se les reconocerán de conformidad con lo dispuesto en la Resolución 01 de febrero del 2012 de la Vicerrectoría Académica y hasta por un máximo de ocho (8) horas semanales.\n\nQue mediante Acuerdo 002 de 2011, se modificó el artículo 43 del Acuerdo 011 de 2002, de manera que se pueden reconocer honorarios a personal pensionado en la modalidad de docentes de hora cátedra por honorarios, hasta por ocho (8) horas semanales.\n\nQue dado que la presente vinculación no genera relación laboral con los docentes, estos se encuentran obligados a realizar los aportes al Sistema Integral de Seguridad Social por su propia cuenta.\n\nQue el Consejo Académico de la Universidad Distrital Francisco José de Caldas, en ejercicio de las funciones establecidas en los literales b) y e) del artículo 18 del Estatuto General (Acuerdo 003 de 1997), mediante Resolución 051 de 28 de septiembre de 2021, aprobó el Calendario Académico del año 2022 para los programas académicos de posgrados.\n\nQue el artículo 1° de la Resolución 051 de 28 de septiembre de 2021 determinó que el inicio de las clases en los programas de posgrado, puede tener lugar en el periodo comprendido entre el 20 de enero y el 19 de febrero de 2022, habiéndose establecido el cierre del periodo para el periodo comprendido entre el 04 de junio y el 24 de junio de 2022.\n\nQue, con base en artículo 1° de la Resolución de Rectoría 210 de 2021, se determina: “Autorizar a las directivas académico – administrativas de la Universidad Distrital Francisco José de Caldas, a dar inicio al plan de retorno seguro y gradual en alternancia de las dependencias administrativas, académico administrativas y comunidad académica, de conformidad con los criterios, condiciones y protocolos de bioseguridad establecidas por las autoridades sanitarias, en el marco de la fase de Aislamiento Selectivo, Distanciamiento Individual Responsable y Reactivación Económica Segura de la pandemia por COVID-19”.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que realicen las actividades lectivas señaladas en el calendario académico y que no son suplidas por los docentes de carrera de la Universidad Distrital.\n\nQue, para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal XXX del XXX de XX enero de 2022.\n\nQue en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "RECONOCIMIENTO. Reconocer honorarios a los siguientes docentes en la modalidad de Hora Cátedra Honorarios (HCH) en los programas de posgrado de la Facultad de XXXXXXXXXXXXXXXXX, conforme al escalafón y dedicación establecidos en la siguiente tabla, para el primer periodo académico de 2022, comprendido entre el XX de XXX de 2022 y hasta el XX de XXX del 2022, con cargo a los recursos asignados en el Certificado de Disponibilidad Presupuestal xxx del xx de enero de 2022 y de conformidad con la Ley 4ª de 1992:\n\n"}
			paragrafo := Paragrafo{Texto: "Los gastos que se ocasionen por la vinculación de los señalados docentes se harán con cargo al Certificado de Disponibilidad Presupuestal (Sueldo Básico) No. XX del XXX de enero de 2022."}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		}
		// accion = "Reconocer Honorarios"
		//nombreDedicacion = "Hora Cátedra Honorarios"
	case "TCO-MTO":
		resolucion = ResolucionCompleta{Consideracion: "Que  el Decreto 1279 de 2002, mediante el cual se establece el régimen salarial y prestacional de los docentes de las universidades estatales, señala en el artículo 3° que “(…) los profesores ocasionales no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto”, precisando que, “no obstante, su vinculación se hace conforme a las reglas que define cada Universidad, con sujeción a lo dispuesto por la Ley 30 de 1992 y demás disposiciones constitucionales y legales vigentes”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial en las modalidades de Tiempo Completo Ocasional (TCO), Medio Tiempo Ocasional (MTO), Hora Cátedra (HC) y Hora Catedra Honorarios (HCH), de que trata el artículo 13 del Acuerdo 011 de 2002 (Estatuto Docente), a término fijo por periodos académicos.\n\nQue mediante Resolución 001 de 15 de febrero de 2012 la Vicerrectoría Académica estableció el procedimiento para la selección y vinculación a la Universidad de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial deberán ser reconocidos en los términos del inciso segundo 2º del artículo 74 de la Ley 30 de 1992, esto es, “mediante resolución”.\n\nQue la Corte Constitucional, en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992 aclaró, entre otras cosas, que la vinculación de los docentes de vinculación especial estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de 18 de diciembre de 2018 del Departamento Administrativo de la Función Pública, “en el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración (…)”.\n\nQue de conformidad con el artículo 2° del Decreto Nacional 447 de 29 de marzo 2022, “(…) a partir del 1° de enero de 2022 se fija el valor del punto salarial para los empleados públicos docentes a quienes se les aplica el Decreto 1279 de 2002 y demás disposiciones que lo modifiquen o adicionen en dieciséis mil cuatrocientos cuarenta y un pesos ($16.441) moneda corriente (…)”.\n\nQue en virtud del mencionado decreto, se acoge y aplica, en lo pertinente, única y expresamente al valor del punto salarial en DIECISÉIS MIL CUATROCIENTOS CUARENTA Y UN PESOS ($16.441) MONEDA CORRIENTE, para los docentes de Vinculación Especial para los programas de pregrado, esto en concordancia con el artículo 2° del Acuerdo 012 de 2002 del Consejo Superior Universitario.\n\nQue en materia de Seguridad y Salud en el Trabajo (SST) para docentes ocasionales de la Universidad Distrital Francisco José de Caldas, se deben informar las obligaciones específicas de que trata el artículo 2.2.4.2.2.16. del Decreto Nacional 1072 de 2015, así como la Resolución de Rectoría 624 de 2017, conforme a los cuales, el docente ocasional, de hora catedra y por honorarios, debe cumplir con las normas del Sistema General de Riesgos Laborales.\n\nQue el Consejo Académico, en uso de las funciones estatutarias establecidas en los literales b) y e) del artículo 18 del Estatuto General -Acuerdo 003 de 1997 del CSU-, mediante Resolución 13 de 15 de marzo de 2022, fijó el calendario académico especial para los periodos académicos 2022-I y 2022-III en los distintos programas de pregrado de la Universidad Distrital Francisco José de Caldas.\n\nQue a través de la Resolución 014 de 29 de marzo de 2022, el Consejo Académico modificó la Resolución 013 de 15 de marzo de 2022, por medio de la cual se fijó el calendario académico especial para los periodos académicos 2022-1 y 2022-lII en los distintos programas académicos de pregrado de la Universidad Distrital Francisco José de Caldas\".\n\nQue la Resolución de Rectoría 087 de 2022, “Por medio de la cual se adoptan las medidas para el regreso del personal administrativo de la Universidad Distrital Francisco José de Caldas a las actividades laborales presenciales”, determina: “El reinicio total de actividades administrativas y académico- administrativas presenciales será a partir del 02 de marzo del 2022. La totalidad de servidores administrativos (libre nombramiento y remoción, carrera administrativa, provisionalidad y temporalidad) deberán asistir diariamente a desempeñar sus funciones de manera presencial en las correspondientes sedes de la Universidad Distrital Francisco José de Caldas”.\n\nQue igualmente, el artículo 4 del mentado acto administrativo señala que “(…) El regreso a las actividades presenciales implica el cumplimiento cabal de la Resolución 777 del 2021 del Ministerio de Salud y Protección Social, así como del Protocolo de Bioseguridad de la Universidad-2022”.\n\nQue se hace necesario garantizar la vinculación de docentes de Tiempo Completo Ocasional o Medio Tiempo Ocasional que realicen las actividades de docencia señaladas en el calendario académico y en el plan de trabajo aprobado y que no son suplidas por los docentes de carrera de la Universidad Distrital.\n\nQue, para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal (Sueldo Básico, Prima, Vacaciones, Prima Navidad, Prima de Servicios y Cesantías) 0xxx del xx de enero de 2022.\n\nEn mérito de lo expuesto,\n\n"}
		articulo = Articulo{Texto: "VINCULACIÓN. Vincular a los siguientes docentes para los periodos académicos 2022-I y 2022-III, comprendidos entre el 26 de abril y el 19 de agosto de 2022, y entre el 30 de agosto de 2022 y el 24 de diciembre de 2022, respectivamente, en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, en la modalidad de Tiempo Completo Ocasional y Medio Tiempo Ocasional, en el escalafón y dedicación establecidas en la siguiente tabla, y con cargo a los recursos asignados en el Certificado de Disponibilidad Presupuestal (Sueldo Básico, Prima, Vacaciones, Prima Navidad, Prima de Servicios y Cesantías) xxx del xx de enero de 2022:\n\n"}
		paragrafo := Paragrafo{Texto: "La vinculación de los docentes referidos en este artículo se suspenderá durante el periodo comprendido entre el 20 y el 29 de agosto de 2022, el cual corresponde a la terminación del periodo académico 2022-I y el inicio del periodo académico 2022-III. Por lo tanto, en ese lapso no existirá para los docentes la obligación de prestar sus servicios, ni para la Universidad la obligación de pagar salarios y prestaciones; pero persistirá la obligación de la institución de efectuar los respectivos aportes a salud y pensión en el porcentaje que le corresponda."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		// accion = "Vincular"
		//nombreDedicacion = "Medio Tiempo Ocasional y Tiempo Completo Ocasional"
	}

	// if tipo != "1" {
	// 	articulo = Articulo{Texto: "Modificar la Resolución No XXX del XXX del XXXX en cuanto al número de horas semanales y el valor total para el " + periodoStr + " Período Académico del " + strconv.Itoa(vigencia) + ", como docentes en la modalidad de " + nombreDedicacion + " de Vinculación Especial, en el escalafón y dedicación establecidas en la siguiente tabla:"}
	// }

	articulos = append(articulos, articulo)

	/////ARTICULO 2
	if dedicacion == "HCH" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "PREPARACIÓN DE CURRICULOS. El proceso de revisión y adaptación de las asignaturas para el uso de herramientas digitales de apoyo, así como la preparación de los diferentes contenidos temáticos y demás actividades académicas mediante la apropiación de dichas herramientas, se harán con base en la distribución horaria establecida por cada facultad.\n\n"}
		paragrafo := Paragrafo{Texto: "De conformidad con los considerados de la Resolución No. 065 del 6 de octubre del .020 del Consejo Académico, los docentes podrán apoyarse en Planestic y demás instancias funcionales de la Universidad, para llevar a cabo el proceso de revisión y adaptación de las asignaturas para el uso de herramientas digitales de apoyo, y la preparación de los diferentes contenidos temáticos."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	if dedicacion == "HCP" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "REMUNERACIÓN. El pago de los servicios prestados por los profesores de vinculación especial a que se refiere el artículo anterior, según su escalafón, se hará previa certificación de las horas efectivamente dictadas, que se encuentren inmersas en el correspondiente plan de trabajo de la gestión académica, expedida por el decano y/o director de proyecto curricular, según corresponda.\n\n"}
		paragrafo := Paragrafo{Texto: "El valor del salario mínimo mensual legal vigente para el reconocimiento y pago de los docentes en cuestión, será el que fije el Gobierno Nacional.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	if dedicacion == "TCO-MTO" {
		//MTO Y TCO
		articulo = Articulo{Texto: "CONDICIONES. El pago de los salarios a los docentes de vinculación especial a que se refiere el artículo anterior, según su escalafón, se hará previa certificación, de las actividades de docencia efectivamente desarrolladas conforme al correspondiente plan de trabajo, expedida por el decano y/o coordinador de proyecto curricular.\n\n"}
		paragrafo := Paragrafo{Texto: "El pago de los salarios a los docentes de vinculación especial se realizará los primeros cinco (5) días hábiles, del mes siguiente a la prestación del servicio.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "El valor del punto salarial en pesos para el reconocimiento y pago de los docentes en cuestión, será el que fije el Gobierno Nacional, mediante decreto, cada año, y que la Universidad Distrital Francisco José de Caldas acoja, mediante acto administrativo.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	if dedicacion == "HCH" && nivel == "PREGRADO" {
		articulo = Articulo{Texto: "CONDICIONES. El pago de los servicios prestados por los profesores de vinculación especial a que se refiere el artículo anterior, según su escalafón, se hará previa certificación de las horas efectivamente dictadas conforme al correspondiente plan de trabajo, expedida por el decano y/o coordinador de proyecto curricular.\n\n"}
		paragrafo := Paragrafo{Texto: "El pago de los salarios a los profesores se realizará los primeros cinco (5) días hábiles, del mes siguiente a la prestación del servicio.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "El valor del punto salarial en pesos para el reconocimiento y pago de los docentes en cuestión será el que fije el Gobierno Nacional mediante decreto cada año y que la Universidad Distrital Francisco José de Caldas acogerá mediante acto administrativo.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	if dedicacion == "HCP" && nivel == "PREGRADO" {
		articulo = Articulo{Texto: "CONDICIONES. El pago de los servicios prestados por los profesores de vinculación especial a que se refiere el artículo anterior, según su escalafón, se hará previa certificación de las horas efectivamente dictadas conforme al correspondiente plan de trabajo, expedida por el decano y/o coordinador de proyecto curricular.\n\n"}
		paragrafo := Paragrafo{Texto: "El pago de los salarios a los profesores se realizará los primeros cinco (5) días hábiles, del mes siguiente a la prestación del servicio.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: ". El valor del punto salarial en pesos para el reconocimiento y pago de los docentes en cuestión será el que fije el Gobierno Nacional mediante decreto cada año y que la Universidad Distrital Francisco José de Caldas acogerá mediante acto administrativo.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}

	articulos = append(articulos, articulo)
	/////ARTICULO 3
	if dedicacion == "HCH" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "OBLIGACIONES. Los docentes a que se refiere el artículo 1° deberán cumplir con las obligaciones inherentes a la naturaleza del servicio, contempladas en la ley, en los reglamentos de la institución y en los planes de trabajo establecidos en el aplicativo de gestión académica, entregados por cada profesor, y aprobados por el coordinador del correspondiente programa académico, decano y/o director, según corresponda.\n\n"}
		paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el Plan de Trabajo del período académico correspondiente y, de manera excepcional, se podrán desarrollar actividades académicas asistidas por las herramientas que ofrecen las Tecnologías de la Información y las Comunicaciones (TICs), garantizando las condiciones de calidad reconocidas en el registro calificado, sin que implique el cambio de modalidad.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "El docente cumplirá con lo establecido en la Circular 046 de 19 de julio de 2017 de Rectoría, con relación al pago de aportes al Sistema Integral de Seguridad Social de forma independiente, de conformidad con la ley.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "El docente deberá cumplir con aquellos procedimientos establecidos en el Sistema Integrado de Gestión de la entidad (SIGUD) para el proceso de Gestión de Docencia, en cuanto le competan.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	if dedicacion == "HCP" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "PREPARACIÓN DE CURRICULOS. El proceso de revisión y adaptación de las asignaturas para el uso de herramientas digitales de apoyo, y la preparación de los diferentes contenidos temáticos y demás actividades académicas, mediante la apropiación de dichas herramientas, se harán con base en la distribución horaria establecida por cada facultad.\n\n"}
		paragrafo := Paragrafo{Texto: "De conformidad con los considerados de la Resolución No. 065 del 6 de octubre del 2020 del Consejo Académico, los docentes podrán apoyarse en Planestic y demás instancias funcionales de la Universidad, para llevar a cabo el proceso de revisión y adaptación de las asignaturas para el uso de herramientas digitales de apoyo, y la preparación de los diferentes contenidos temáticos."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	if dedicacion == "TCO-MTO" {
		//MTO Y TCO
		articulo = Articulo{Texto: "OBLIGACIONES. Los docentes en cuestión deberán cumplir con las obligaciones inherentes a la naturaleza del servicio contempladas en la ley, en los reglamentos de la Universidad Distrital Francisco José de Caldas y en los planes de trabajo establecidos en el aplicativo de gestión académica, entregados por cada docente, y aprobados por el coordinador del correspondiente proyecto curricular y/o decano.\n\n"}
		paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el respectivo plan de trabajo.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	if dedicacion == "HCP" && nivel == "PREGRADO" {
		articulo = Articulo{Texto: "OBLIGACIONES. Los docentes en cuestión deberán cumplir con las obligaciones inherentes a la naturaleza del servicio público de educación superior contempladas en la ley, en los reglamentos de la institución y en los planes de trabajo establecidos en el aplicativo de gestión académica, entregados por cada docente, y aprobados por el coordinador del correspondiente proyecto curricular y/o decano.\n\n"}
		paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el plan de trabajo.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}

	if dedicacion == "HCH" && nivel == "PREGRADO" {

		// Pregrado y HCH
		articulo = Articulo{Texto: "OBLIGACIONES. Los docentes en cuestión deberán cumplir con las obligaciones inherentes a la naturaleza del servicio público de educación superior contempladas en la ley, en los reglamentos de la institución y en los planes de trabajo establecidos en el aplicativo de gestión académica, entregados por cada docente, y aprobados por el coordinador del correspondiente proyecto curricular y/o decano.\n\n"}
		paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el plan de trabajo.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	articulos = append(articulos, articulo)
	////ARTICULO 4
	if dedicacion == "HCH" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "TERMINACIÓN ANTICIPADA. En caso de incumplimiento o retiro del docente, la Universidad Distrital Francisco José de Caldas, a través del correspondiente ordenador del gasto, mediante acto administrativo motivado, previo trámite en el cual se garantice al docente el ejercicio de los derechos de contradicción y de defensa, declarará la terminación del vínculo y su liquidación con corte a la fecha del respectivo cumplido, pagando lo correspondiente, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
		articulos = append(articulos, articulo)
	}
	if dedicacion == "HCP" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "OBLIGACIONES. Los docentes en cuestión deberán cumplir con las obligaciones inherentes a la naturaleza del servicio, contempladas en la ley, en los reglamentos de la Universidad Distrital y en los planes de trabajo establecidos en el aplicativo de gestión académica, entregados por cada profesor, y aprobados por el coordinador del correspondiente programa académico, decano y/o director, según corresponda.\n\n"}
		paragrafo := Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el Plan de Trabajo del período académico correspondiente y, de manera excepcional, se podrán desarrollar actividades académicas asistidas por las herramientas que ofrecen las Tecnologías de la Información y las Comunicaciones (TICs), garantizando las condiciones de calidad reconocidas en el registro calificado, sin que esto implique cambio de modalidad.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "El docente deberá cumplir con los procedimientos establecidos en el Sistema Integrado de Gestión de la entidad (SIGUD) para el proceso de Gestión de docencia, en cuanto le competan.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		articulos = append(articulos, articulo)

	} else {
		if dedicacion == "HCP" && nivel == "PREGRADO" {
			articulo = Articulo{Texto: "TERMINACIÓN. - En caso de incumplimiento, retiro del docente o en el momento que se declare la cancelación de actividades académicas, por parte del Consejo Superior Universitario de la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
			articulos = append(articulos, articulo)
		}
		if dedicacion == "HCH" && nivel == "PREGRADO" {
			// Pregrado y HCH
			articulo = Articulo{Texto: "TERMINACIÓN. - En caso de incumplimiento, retiro del docente o en el momento que se declare la cancelación de actividades académicas, por parte del Consejo Superior Universitario de la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
			articulos = append(articulos, articulo)
		}
		if dedicacion == "TCO-MTO" {
			articulo = Articulo{Texto: ". TERMINACIÓN. En caso de incumplimiento, retiro del docente o en el momento que se declare la cancelación de actividades académicas por parte del Consejo Superior Universitario la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
			articulos = append(articulos, articulo)
		}
	}
	/////ARTICULO 5
	if dedicacion == "HCH" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "IMPUTACIÓN PRESUPUESTAL. El gasto que ocasione el presente acto administrativo se hará con cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal.\n\n"}
		paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias realizadas por la Secretaría de Hacienda Distrital, y se realizarán dentro de los primeros cinco (5) días hábiles de cada mes."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	} else {
		if dedicacion == "HCH" && nivel == "PREGRADO" {
			// Pregrado y HCH
			articulo = Articulo{Texto: "AMPARO PRESUPUESTAL. El gasto que ocasione el presente acto administrativo se hará con cargo al presupuesto de la actual vigencia fiscal, previa certificación de disponibilidad presupuestal correspondiente.\n\n"}
			paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias que deba efectuar la Secretaría de Hacienda Distrital.\n\n"}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		}
		if dedicacion == "HCP" && nivel == "POSGRADO" {
			articulo = Articulo{Texto: "TERMINACIÓN ANTICIPADA. En caso de incumplimiento o retiro del docente, la Universidad Distrital Francisco José de Caldas, a través del correspondiente ordenador del gasto, mediante acto administrativo motivado, previo trámite en el cual se garantice al docente el ejercicio de los derechos de contradicción y de defensa, declarará la terminación del vínculo y la liquidación con corte a la fecha del respectivo cumplido, pagando lo correspondiente, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
		}
		if dedicacion == "HCP" && nivel == "PREGRADO" {
			//HCP Pregrado
			articulo = Articulo{Texto: "AMPARO PRESUPUESTAL. El gasto que ocasione el presente acto administrativo se hará con cargo al presupuesto de la actual vigencia fiscal, previa certificación de disponibilidad presupuestal correspondiente.\n\n"}
			paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias que deba efectuar la Secretaría de Hacienda Distrital.\n\n"}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		}
		if dedicacion == "TCO-MTO" {
			//tco-mto
			articulo = Articulo{Texto: "AMPARO PRESUPUESTAL. El gasto que ocasione el presente acto administrativo se hará con cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal.\n\n"}
			paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias que deba efectuar la Secretaría de Hacienda Distrital.\n\n"}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		}
	}
	articulos = append(articulos, articulo)
	/////ARTICULO 6
	if dedicacion == "HCH" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "SUSPENSIÓN. En el supuesto de que se declare la suspensión de actividades académicas por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial la obligación de prestar sus servicios; y para la Universidad la de pagar los honorarios correspondientes al periodo suspendido.\n\nSi la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, proferido por el ordenador del gasto, podrá declarar la terminación del vínculo y efectuar la liquidación correspondiente, con corte a la fecha del respectivo cumplido, pagando lo correspondiente, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
		articulos = append(articulos, articulo)
	}
	if dedicacion == "HCP" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "IMPUTACIÓN PRESUPUESTAL. El gasto que ocasione el presente acto administrativo, se hará con cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal.\n\n"}
		paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias realizadas por la Secretaría de Hacienda Distrital y se realizarán dentro de los primeros cinco (5) días hábiles de cada mes.\n\n"}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	} else {
		if dedicacion == "HCH" && nivel == "PREGRADO" {
			// Pregrado y HCH
			articulo = Articulo{Texto: "SUSPENSIÓN. En el supuesto de que se declare la suspensión de actividades académicas, por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial, la obligación de prestar sus servicios y para la Universidad, la de pagar los salarios y prestaciones correspondientes al periodo suspendido, pero persistirá, para esta última, la de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda. En este caso, la fecha de terminación de la vinculación se prolongará por un período igual al de la suspensión\n\nSi la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, podrá declarar la terminación del vínculo y efectuar la liquidación correspondiente, con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
			articulos = append(articulos, articulo)
		}
		if dedicacion == "HCP" && nivel == "PREGRADO" {
			//HCP
			articulo = Articulo{Texto: "SUSPENSIÓN. En el supuesto de que se declare la suspensión de actividades académicas, por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial, la obligación de prestar sus servicios y para la Universidad, la de pagar los salarios y prestaciones correspondientes al periodo suspendido, pero persistirá, para esta última, la de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda. En este caso, la fecha de terminación de la vinculación se prolongará por un período igual al de la suspensión\n\nSi la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, podrá declarar la terminación del vínculo y efectuar la liquidación correspondiente, con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
			articulos = append(articulos, articulo)
		}
		if dedicacion == "TCO-MTO" {
			//TCO-MTO
			articulo = Articulo{Texto: "MODIFICACIÓN DEL PLAN DE TRABAJO. Los docentes de Tiempo Completo Ocasional y Medio Tiempo Ocasional que no desarrollen horas lectivas por situaciones que impidan el normal desarrollo del calendario académico, deberán modificar su plan de trabajo respecto de las horas lectivas afectadas, con el fin de desarrollar otras actividades, agotando para el efecto el mismo procedimiento utilizado para la aprobación inicial del mismo.\n\n"}
			paragrafo := Paragrafo{Texto: "Si excepcionalmente y por cualquier circunstancia no es posible modificar el plan de trabajo, la vinculación se suspenderá con fundamento en las situaciones referidas en el inciso anterior, y la fecha de terminación de la misma se prolongará automáticamente por un tiempo igual al de la suspensión.\n\n"}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
			paragrafo = Paragrafo{Texto: "Durante la suspensión, cesarán correlativamente los efectos salariales y prestacionales correspondientes, pero persistirá la obligación de la institución de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda.\n\n"}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
			paragrafo = Paragrafo{Texto: "En los casos mencionados en el inciso anterior, si la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad, mediante acto administrativo motivado, podrá declarar la terminación del vínculo y efectuar la liquidación correspondiente con corte a la fecha del respectivo cumplido, pagando las correspondientes prestaciones sociales.\n\n"}
			articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
			articulos = append(articulos, articulo)
		}
	}
	/////ARTICULO 7
	if dedicacion == "HCH" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "OBLIGACIONES EN MATERIA DE RIESGOS LABORALES. Los docentes a que se refiere el presente acto administrativo, en los términos previstos en el artículo 2.2.4.2.2.16. del Decreto Nacional 1072 de 2015 y la Resolución de Rectoría 624 de 2017, darán cumplimiento a las normas del Sistema General de Riesgos Laborales.\n\n"}
		articulos = append(articulos, articulo)
	}
	if dedicacion == "HCP" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "SUSPENSIÓN. En el supuesto de que se declare la suspensión de actividades académicas, por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial, la obligación de prestar sus servicios y para la Universidad, la de pagar los salarios y prestaciones correspondientes al periodo suspendido, pero persistirá, para esta última, la de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda. En este caso, la fecha de terminación de la vinculación se prolongará por un período igual al de la suspensión\n\nSi la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, podrá declarar la terminación del vínculo y efectuar la liquidación correspondiente, con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos.\n\n"}
		articulos = append(articulos, articulo)

	} else {
		if dedicacion == "TCO-MTO" {
			articulo = Articulo{Texto: "RIESGOS LABORALES. Los docentes ocasionales a que se refiere el presente acto administrativo, en los términos previstos en el artículo 2.2.4.2.2.16. del Decreto Nacional 1072 de 2015 y la Resolución de Rectoría 624 de 2017, darán cumplimiento a las normas del Sistema General de Riesgos Laborales.\n\n"}
			articulos = append(articulos, articulo)
		} else {
			if dedicacion == "HCH" && nivel == "PREGRADO" {
				articulo = Articulo{Texto: "RIESGOS LABORALES. Los docentes hora cátedra a que se refiere el presente acto administrativo, en los términos previstos en el artículo 2.2.4.2.2.16. del Decreto Nacional 1072 de 2015 y la Resolución de Rectoría 624 de 2017, darán cumplimiento a las normas del Sistema General de Riesgos Laborales.\n\n"}
				articulos = append(articulos, articulo)
			}
			if dedicacion == "HCP" && nivel == "PREGRADO" {
				// Pregrado y HCP
				articulo = Articulo{Texto: "RIESGOS LABORALES. Los docentes hora cátedra a que se refiere el presente acto administrativo, en los términos previstos en el artículo 2.2.4.2.2.16. del Decreto Nacional 1072 de 2015 y la Resolución de Rectoría 624 de 2017, darán cumplimiento a las normas del Sistema General de Riesgos Laborales.\n\n"}
				articulos = append(articulos, articulo)
			}
		}
	}
	/////ARTICULO 8
	if dedicacion == "HCH" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "INHABILIDAD O INCOMPATIBILIDAD. Comuníquese la presente resolución a los docentes mencionados en el artículo 1°, quienes deberán manifestar, bajo la gravedad de juramento, el cual se entenderá prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, que no tienen cruces de horarios y que el monto de los honorarios que perciben por concepto de hora cátedra no corresponden a más de ocho (8) horas diarias de trabajo a varias entidades.\n\n"}
		articulos = append(articulos, articulo)
	}
	if dedicacion == "HCP" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "RIESGOS LABORALES. Los docentes hora cátedra a que se refiere el presente acto administrativo, en los términos previstos en el artículo 2.2.4.2.2.16. del Decreto 1072 de 2015 y la Resolución de Rectoría No. 624 de 2017, darán cumplimiento a las normas del Sistema General de Riesgos Laborales.\n\n"}
		articulos = append(articulos, articulo)
	} else {
		if dedicacion == "HCH" && nivel == "PREGRADO" {
			// Pregrado y HCH
			articulo = Articulo{Texto: "INHABILIDAD O INCOMPATIBILIDAD. Comuníquese la presente resolución a los docentes mencionados en el artículo 1º, quienes deberán manifestar, bajo la gravedad de juramento, que se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, y que no tienen cruces de horarios.\n\n"}
			articulos = append(articulos, articulo)
		} else {
			if dedicacion == "TCO-MTO" {
				articulo = Articulo{Texto: "INHABIBILIDAD O INCOMPATIBILIDAD. Comuníquese la presente resolución a los docentes mencionados en el artículo primero, quienes deberán manifestar, bajo la gravedad de juramento, el cual se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad, establecida en las normas pertinentes y aplicables, así como que no tienen cruces de horarios.\n\n"}
				articulos = append(articulos, articulo)
			}
			if dedicacion == "HCP" && nivel == "PREGRADO" {
				//HCPPregrado y Posgrado
				articulo = Articulo{Texto: "INHABILIDAD O INCOMPATIBILIDAD. Comuníquese la presente resolución a los docentes mencionados en el artículo 1º, quienes deberán manifestar, bajo la gravedad de juramento, que se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, y que no tienen cruces de horarios.\n\n"}
				articulos = append(articulos, articulo)
			}
		}
	}
	/////ARTICULO 9
	// Honorarios
	if dedicacion == "HCH" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "VIGENCIA. El presente acto administrativo surte efectos de conformidad con la Resolución 051 de 28 de septiembre de 2021, por medio de la cual se expide el calendario académico para el 2022, en concreto, para las actividades académicas relacionadas con el primer periodo académico de 2022, el cual va del 20 de enero de 2022 y hasta el 24 de junio 2022.\n\n"}
		articulos = append(articulos, articulo)
	}
	if dedicacion == "HCH" && nivel == "PREGRADO" {
		articulo = Articulo{Texto: "VIGENCIA. El presente acto administrativo surte efectos de conformidad con la Resolución 014 de 29 de marzo de 2022 “por medio de la cual se modificó el calendario académico especial para los periodos académicos 2022-I y 2022-III”, y demás normas concordantes, complementarias y modificatorias.\n\n"}
		articulos = append(articulos, articulo)
	}
	//TCO-MTO
	if dedicacion == "TCO-MTO" {
		articulo = Articulo{Texto: "VIGENCIA. El presente acto administrativo surte efectos de conformidad con la Resolución 014 de 29 de marzo de 2022, por medio de la cual se modificó el calendario académico especial para los periodos académicos 2022-I y 2022-III, y demás normas concordantes, complementarias y modificatorias.\n\n"}
		articulos = append(articulos, articulo)
	}
	// Prestación
	if dedicacion == "HCP" && nivel == "PREGRADO" {
		articulo = Articulo{Texto: "VIGENCIA. El presente acto administrativo surte efectos de conformidad con la Resolución 014 de 29 de marzo de 2022 “por medio de la cual se modificó el calendario académico especial para los periodos académicos 2022-I y 2022-III”, y demás normas concordantes, complementarias y modificatorias.\n\n"}
		articulos = append(articulos, articulo)
	}
	if dedicacion == "HCP" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "DECLARACIÓN DE AUSENCIA DE INHABILIDADES E INCOMPATIBILIDADES. Comuníquese la presente resolución a los docentes mencionados en el artículo primero, quienes deberán manifestar, bajo la gravedad de juramento que se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, que no tienen cruces de horarios.\n\n"}
		articulos = append(articulos, articulo)

	}

	/////ARTICULO 10
	if dedicacion == "HCP" && nivel == "POSGRADO" {
		articulo = Articulo{Texto: "VIGENCIA. -  El presente acto administrativo surte efectos de conformidad con la Resolución 051 de 28 de septiembre 2021, por medio del cual se expide el Calendario Académico para el año 2022, en concreto, para las actividades académicas relacionadas con el primer periodo académico del año 2022, que corresponde del 20 de enero de 2022 y hasta el 24 de junio de 2022.\n\n"}
		articulos = append(articulos, articulo)
	}

	resolucion.Articulos = articulos

	return resolucion
}
