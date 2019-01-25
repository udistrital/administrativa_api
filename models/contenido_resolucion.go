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

func GetTemplateResolucion(dedicacion, nivel, periodo, tipo string) (res ResolucionCompleta) {
	var resolucion ResolucionCompleta
	var articulos []Articulo
	var articulo Articulo
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
			resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de Noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\nQue mediante Resolución Nro. 001 del 15 de Febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial, hora cátedra, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos, en los términos del artículo 73 de la Ley 30 de 1992.\n\nQue el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue mediante Resolución Nro. 095 de octubre 30 de 2018, el Consejo Académico, optó por “SUSPENDER, a partir del 1º de noviembre de 2018 de manera temporal, el Calendario Académico contenido en la Resolución No. 182 de diciembre 19 de 2017…, hasta tanto se permita el acceso a las instalaciones para el normal desarrollo de las actividades académicas en la Institución”.\n\nQue mediante Resolución 005 de enero 22 de 2019, se modificó la Resolución 110 de 2018 y 01 de 2019, estableciéndose el calendario para el reinicio de las actividades académicas, para finalizar el período académico 2018-3, entre el 28 de enero de 2019 y el 9 de marzo de 2019.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que aseguren la continuidad de las actividades lectivas señaladas en el calendario académico y la culminación satisfactoria del periodo académico 2018-III.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. XX del XX de enero de 2019.\n\nQue, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "Vincular para reanudar las actividades académicas en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, tendientes a finalizar el periodo académico de 2018-III, como docentes en la modalidad de Hora Cátedra de Vinculación Especial, en el escalafón y dedicación establecidas en la tabla, para el periodo comprendido entre el 28 de enero de 2019 y el 9 de marzo de 2019, a los siguientes profesores:"}
		}
		if nivel == "POSGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de Noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\nQue mediante Resolución Nro. 001 del 15 de Febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial, hora cátedra, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos, en los términos del artículo 73 de la Ley 30 de 1992.\n\nQue el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones: 317 de septiembre 8 de 2006 y 318 de Septiembre 8 de 2006, Acuerdos 005 y 007 de 2001, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y Ley 4 de 1992).\n\nQue mediante Acuerdo 006 de julio 19 de 2002 se fija el valor de la Hora Cátedra por Honorarios y se establece un número máximo de horas para los docentes de carrera que presten servicios a la Universidad Distrital Francisco José de Caldas, en los programas de Postgrado, y modifican parcialmente los Acuerdos 005 y 007 de 2001.\n\nQue mediante Acuerdo 002 de enero 31 de 2003 se modifica y reglamenta el Acuerdo 001 de enero 17 de 2003.\n\nQue mediante Acuerdo 005 de julio 27 de 2001 se establece que por medio del cual se fija la Hora Cátedra y se establece el número máximo de horas, para los docentes que presten servicios a la Universidad Distrital en los programas de Postgrado.\n\nQue el Consejo Académico mediante Resolución 112 de Diciembre 18 de 2018  emitió el calendario académico para las actividades del año 2019 a desarrollarse en los programas de posgrado en la Universidad, para este caso se reconocerán de conformidad con los módulos de las horas efectivamente dictadas y aprobadas en los diferentes programas como carga lectiva.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que aseguren las actividades lectivas señaladas en el calendario académico 2019-I.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. XX del XX de enero de 2019.\n\nQue, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "Vincular para el " + periodoStr + " Periodo Académico del año " + strconv.Itoa(vigencia) + ", docentes de Vinculación Especial en el escalafón y dedicación establecidas en la tabla, a los siguientes docentes:"}
		}
		// accion = "Vincular"
		nombreDedicacion = "Hora Cátedra"
	case "HCH":
		if nivel == "PREGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de Noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\nQue mediante Resolución Nro. 001 del 15 de Febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial, hora cátedra, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos, en los términos del artículo 73 de la Ley 30 de 1992.\n\nQue el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de enero 31 de 2003, 0013-A de enero 31 de 2003, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\nQue conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos del reconocimiento de honorarios el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue mediante Resolución Nro. 095 de octubre 30 de 2018, el Consejo Académico, optó por “SUSPENDER, a partir del 1º de noviembre de 2018 de manera temporal, el Calendario Académico contenido en la Resolución No. 182 de diciembre 19 de 2017…, hasta tanto se permita el acceso a las instalaciones para el normal desarrollo de las actividades académicas en la Institución”.\n\nQue mediante Resolución 005 de enero 22 de 2019, se modificó la Resolución 110 de 2018 y 01 de 2019, estableciéndose el calendario para el reinicio de las actividades académicas, para finalizar el período académico 2018-3, entre el 28 de enero de 2019 y el 9 de marzo de 2019.\n\nQue se hace necesario garantizar el reconocimiento de honorarios a docentes hora cátedra Honorarios que aseguren la continuidad de las actividades lectivas señaladas en el calendario académico y la culminación satisfactoria del periodo académico 2018-III.\n\nQue de conformidad con el artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018 establece “A partir del 1° de enero de 2018, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en trece mil quinientos noventa y ocho pesos ($13.598) moneda corriente”.\n\nQue mediante Resolución 072 del quince de marzo del 2018 emitida por Rectoría mediante la cual acoge únicamente el valor del punto para los docentes de vinculación especial a quienes se les vinculó o reconoció honorarios en esta vigencia con el valor del punto anterior, y se hace necesario dar cumplimiento al artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018.\n\nQue mediante Acuerdo 002 de marzo 17 de 2011 se modifica el artículo 43 del Acuerdo 011 de 2002, por el cual se pueden reconocer honorarios personas pensionadas en la modalidad de docentes de hora- cátedra por honorarios.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. XX del XX de enero de 2019.\n\nQue, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "Reconocer Honorarios para reanudar las actividades académicas en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, tendientes a finalizar el periodo académico de 2018-III, como docentes en la modalidad de Hora Cátedra Honorarios de Vinculación Especial, en el escalafón y dedicación establecidas en la tabla, para el periodo comprendido entre el 28 de enero de 2019 y el 9 de marzo de 2019, a los siguientes profesores:"}
		}
		if nivel == "POSGRADO" {
			resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto en las disposiciones constitucionales y legales”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de Noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\nQue mediante Resolución Nro. 001 del 15 de Febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial, hora cátedra, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos, en los términos del artículo 73 de la Ley 30 de 1992.\n\nQue el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 73 de la Ley 30 de 1992, aclaró, entre otras cosas, que, la vinculación de los docentes de hora cátedra estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones: 317 de septiembre 8 de 2006 y 318 de Septiembre 8 de 2006, Acuerdos 005 y 007 de 2001, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y Ley 4 de 1992).\n\nQue mediante Acuerdo 006 de julio 19 de 2002 se fija el valor de la Hora Cátedra por Honorarios y se establece un número máximo de horas para los docentes de carrera que presten servicios a la Universidad Distrital Francisco José de Caldas, en los programas de Postgrado, y modifican parcialmente los Acuerdos 005 y 007 de 2001.\n\nQue mediante Acuerdo 002 de enero 31 de 2003 se modifica y reglamenta el Acuerdo 001 de enero 17 de 2003.\n\nQue mediante Acuerdo 005 de julio 27 de 2001 se establece que por medio del cual se fija la Hora Cátedra y se establece el número máximo de horas, para los docentes que presten servicios a la Universidad Distrital en los programas de Postgrado.\n\nQue el Consejo Académico mediante Resolución 112 de Diciembre 18 de 2018  emitió el calendario académico para las actividades del año 2019 a desarrollarse en los programas de posgrado en la Universidad, para este caso se reconocerán de conformidad con los módulos de las horas efectivamente dictadas y aprobadas en los diferentes programas como carga lectiva.\n\nQue se hace necesario garantizar la vinculación de docentes hora cátedra que aseguren las actividades lectivas señaladas en el calendario académico 2019-I.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. XX del XX de enero de 2019.\n\nQue, en mérito de lo expuesto,\n\n"}
			articulo = Articulo{Texto: "Reconocer Honorarios para el " + periodoStr + " Periodo Académico del año " + strconv.Itoa(vigencia) + ", docentes de Vinculación Especial en el escalafón y dedicación establecidas en la tabla, a los siguientes docentes:"}
		}
		// accion = "Reconocer Honorarios"
		nombreDedicacion = "Hora Cátedra Honorarios"
	case "TCO-MTO":
		resolucion = ResolucionCompleta{Consideracion: "Que mediante los artículos 3º y 4º del Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció la naturaleza jurídica, laboral y prestacional de los docentes de las universidades estatales u oficiales del orden nacional, departamental, municipal y distrital, prescribiendo en su artículo 3º, que “[l]os profesores de hora-cátedra de las Universidades estatales u oficiales distintas a la Universidad Nacional de Colombia no son empleados públicos docentes de régimen especial ni pertenecen a la carrera profesoral y, por consiguiente, sus condiciones salariales y prestacionales no están regidas por el presente Decreto, sino por las reglas contractuales que en cada caso se convengan, conforme a las normas internas de cada Universidad, con sujeción a lo dispuesto por la ley 30 de 1992 y demás disposiciones constitucionales y legales vigentes”.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes de vinculación especial, en las modalidades de hora cátedra (HC), medio tiempo ocasional (MTO) y tiempo completo ocasional (TCO), en virtud de lo consagrado en el artículo 13 del Acuerdo 011 de Noviembre 15 de 2002 (Estatuto Docente), a término fijo, por periodos académicos.\n\nQue mediante Resolución Nro. 001 del 15 de Febrero de 2012, proferida por la Vicerrectoría Académica, se establece el procedimiento para la selección y vinculación a la Universidad, de docentes de vinculación especial.\n\nQue los servicios de los docentes de vinculación especial, ocasionales de medio tiempo y tiempo completo, de que trata el Acuerdo 011 de noviembre 15 de 2002, deberán ser reconocidos mediante resolución, en los términos del inciso segundo del artículo 74 de la Ley 30 de 1992.\n\nQue el artículo 128 de la Constitución Política de Colombia, establece que nadie podrá desempeñar simultáneamente más de un empleo público, ni recibir más de una asignación que provenga del tesoro público, salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4ª de 1992 determinó, como excepción a dicha regla, entre otras, los honorarios percibidos por concepto de hora cátedra.\n\nQue la Corte Constitucional en Sentencia C-006 de 1996, a través de la cual se declararon inexequibles algunos apartes del artículo 74 de la Ley 30 de 1992, aclaró, entre otras cosas, que, contrariamente al régimen prestacional docente consagrado en el Decreto 1279 del 2002, la vinculación de los docentes de Medio Tiempo y Tiempo Completo ocasional estará mediada por su transitoriedad o temporalidad, de lo cual se deriva la proporcionalidad prestacional al tiempo laborado.\n\nQue conforme al concepto de fecha diciembre 18 de 2018, emitido por el Departamento Administrativo de la Función Pública, “[e]n el caso de la Universidad Distrital Francisco José de Caldas, debe acogerse  a lo dispuesto en sus estatutos internos o en el acuerdo en el cual defina sus políticas y criterios generales para regular los procesos de selección, vinculación y contratación de los docentes ocasionales y catedráticos, en especial en lo relacionado con el reconocimiento y pago de prestaciones sociales y remuneración…”.\n\nQue conforme al parágrafo 1º del artículo quinto de la Resolución Nro. 001 de 2012, “[p]ara efectos de pago y liquidación el mes comprenderá (4) semanas o (30) días laborales en Proyectos Académicos de Pregrado y en los espacios académicos de Posgrado se pagará de acuerdo al número de horas efectivamente dictadas en cada período”.\n\nQue mediante Resolución Nro. 095 de octubre 30 de 2018, el Consejo Académico, optó por “SUSPENDER, a partir del 1º de noviembre de 2018 de manera temporal, el Calendario Académico contenido en la Resolución No. 182 de diciembre 19 de 2017…, hasta tanto se permita el acceso a las instalaciones para el normal desarrollo de las actividades académicas en la Institución”.\n\nQue mediante Resolución Nro. 005 de enero 22 de 2019, se modificó la Resolución 110 de 2018 y 01 de 2019, estableciéndose el calendario para el reinicio de las actividades académicas, para finalizar el período académico 2018-3, entre el 28 de enero de 2019 y el 9 de marzo de 2019.\n\nQue según lo establecido en el artículo 1º de la mencionada Resolución Nro. 110 de diciembre 13 de 2018 del Consejo Académico, modificada mediante Resolución Nro. 05 de 2019, se hace necesario vincular los docentes de vinculación especial de medio tiempo y tiempo completo ocasional, que se requieren para el desarrollo de las actividades académicas en cuestión.\n\nQue de conformidad con el artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018 establece “A partir del 1° de enero de 2018, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en trece mil quinientos noventa y ocho pesos ($13.598) moneda corriente”.\n\nQue mediante Resolución 072 del quince de marzo del 2018 emitida por Rectoría mediante la cual acoge únicamente el valor del punto para los docentes de vinculación especial a quienes se les vinculó o reconoció honorarios en esta vigencia con el valor del punto anterior, y se hace necesario dar cumplimiento al artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018.\n\nQue el Consejo Superior Universitario mediante Resolución Nro. 02 de 17 de enero de 2019, determinó excepcionar única y exclusivamente, para el año 2019, lo dispuesto en el Acuerdo Nro. 01 de 2018 del Consejo Superior Universitario, en lo relacionado con el límite de diez (10) meses de vinculación de los docentes ocasionales de Medio Tiempo y Tiempo Completo y por ende, autorizar a la Vicerrectoría Académica y a las Decanaturas de Facultad, para que según viabilidad jurídica y tomando como prioridad los más altos valores de la Universidad y el equilibrio financiero y presupuestal de la misma, vinculen a los profesores ocasionales de Medio Tiempo y Tiempo Completo de forma excepcional, hasta por once (11) meses, según las necesidades del servicio y bajo su responsabilidad.\n\nQue se hace necesario garantizar la vinculación de docentes que aseguren la continuidad de las actividades lectivas señaladas en el calendario académico y la culminación satisfactoria de los periodos académicos 2018-III, 2019-I y 2019-III.\n\nQue para efectos presupuestales, el presente acto administrativo se ejecutará con cargo a los recursos de que trata el Certificado de Disponibilidad Presupuestal No. XX del XX de enero de 2019.\n\nQue, en mérito de lo expuesto,\n\n"}
		articulo = Articulo{Texto: "Vincular, para reanudar las actividades académicas en los programas de pregrado de la Universidad Distrital Francisco José de Caldas, tendientes a finalizar el período académico de 2018-III, así como para desarrollar lo relacionado con los períodos académicos 2019-I y 2019-III, a partir del 28 de enero de 2019 y hasta el 13 de diciembre de 2019, como docentes en la modalidad de Medio Tiempo y Tiempo Completo Ocasional de Vinculación Especial, en el escalafón y dedicación establecidas en la tabla, a los siguientes profesores:"}
		// accion = "Vincular"
		nombreDedicacion = "Medio Tiempo Ocasional y Tiempo Completo Ocasional"
	}

	if tipo != "1" {
		articulo = Articulo{Texto: "Modificar la Resolución No XXX del XXX del XXXX en cuanto al número de horas semanales y el valor total para el " + periodoStr + " Período Académico del " + strconv.Itoa(vigencia) + ", como docentes en la modalidad de " + nombreDedicacion + " de Vinculación Especial, en el escalafón y dedicación establecidas en la siguiente tabla:"}
	}

	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "El pago de los servicios prestados por los profesores de vinculación especial a que se refiere el artículo anterior, según su escalafón, se hará  previa certificación de las horas efectivamente dictadas, que se encuentren inmersas en el correspondiente plan de trabajo de la gestión académica, expedida por el decano y/o director de proyecto curricular."}
	paragrafo := Paragrafo{Texto: "El valor del punto en pesos para el reconocimiento y pago de los docentes de hora cátedra, será el que fije el Gobierno Nacional mediante decreto, cada año, y que la Universidad Distrital Francisco José de Caldas acoja mediante acto administrativo, respecto de los docentes de vinculación especial."}
	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "Los docentes en cuestión deberán cumplir con las obligaciones inherentes a la naturaleza del servicio, contempladas en la ley, en los reglamentos de la Universidad Distrital y en los planes de trabajo establecidos en el aplicativo de gestión académica, entregados por cada profesor, y aprobados por el coordinador del correspondiente programa académico, decano y/o director."}
	if dedicacion == "TCO-MTO" {
		paragrafo = Paragrafo{Texto: "Los Planes de Trabajo objeto de la presente vinculación, serán acordados o entregados, para cada periodo académico, con base en el calendario que para tal efecto, expida el Consejo Académico de la Universidad Distrital Francisco José de Caldas."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el Plan de Trabajo para cada período académico."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	} else {
		paragrafo = Paragrafo{Texto: "Las funciones propias de la vinculación docente serán aquellas que se fijen en el Plan de Trabajo del período académico correspondiente."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	articulos = append(articulos, articulo)
	if dedicacion == "HCP" {
		articulo = Articulo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos."}
		articulos = append(articulos, articulo)
	}
	if dedicacion == "HCH" {
		articulo = Articulo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad Distrital Francisco José de Caldas, mediante acto administrativo motivado, declarará la terminación del reconocimiento con corte a la fecha del respectivo cumplido, cancelando los correspondientes Honorarios, conforme al cálculo que efectúe la División de Recursos Humanos."}
		articulos = append(articulos, articulo)
	}
	articulo = Articulo{Texto: "El gasto que ocasione el presente acto administrativo, se hará con cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal."}
	paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias de la Secretaría de Hacienda Distrital."}
	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	articulos = append(articulos, articulo)
	if dedicacion == "HCH" {
		articulo = Articulo{Texto: "Bajo el supuesto que, se declare la suspensión de actividades académicas por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial la obligación de prestar sus servicios y para la Universidad, la de reconocer los honorarios correspondientes al periodo suspendido. Si la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas mediante acto administrativo motivado, podrá declarar la terminación con corte a la fecha del respectivo cumplido, cancelando honorarios correspondientes, conforme al cálculo que efectúe la División de Recursos Humanos."}
		articulos = append(articulos, articulo)
	} else {
		articulo = Articulo{Texto: "Bajo el supuesto que, se declare la suspensión de actividades académicas por parte de los órganos competentes de la Universidad Distrital Francisco José de Caldas, cesará automáticamente para el docente de vinculación especial la obligación de prestar sus servicios y para la Universidad, la de pagar los salarios y prestaciones correspondientes al periodo suspendido, pero persistirá para esta última, la de efectuar los respectivos aportes a salud y pensión, en el porcentaje que le corresponda. Si la suspensión de actividades académicas persiste por más de quince (15) días calendario, la Universidad Distrital Francisco José de Caldas mediante acto administrativo motivado, podrá declarar la terminación del vínculo laboral y la liquidación con corte a la fecha del respectivo cumplido, cancelando las correspondientes prestaciones sociales, conforme al cálculo que efectúe la División de Recursos Humanos."}
		articulos = append(articulos, articulo)
	}
	articulo = Articulo{Texto: "Comuníquese la presente resolución a los docentes mencionados en el artículo 1º, quienes deberán manifestar, bajo la gravedad de juramento que se entiende prestado con la aprobación del formato de carga en el aplicativo de gestión académica, que no se encuentran incursos en causal de inhabilidad o incompatibilidad establecida en las normas pertinentes y aplicables, que no tienen cruces de horarios ni ostentan otra vinculación de carácter público, diferente a hora cátedra, en entidades de educación oficiales, siempre y cuando los honorarios no sumen más de ocho (8) horas diarias de trabajo a varias entidades."}
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "El presente acto administrativo se expide a los XXXXXX (XX) días del mes de XXXXXX del año " + strconv.Itoa(vigencia) + " y surte efectos en los términos de la. Resolución Nro. 110 de diciembre 13 de 2018, modificada por la Resolución Nro. 05 de enero 22 de 2019 del Consejo Académico, es decir, dependiendo de la normalización de las actividades, tanto administrativas como académicas en la Facultad, para el " + periodoStr + " Periodo Académico del año " + strconv.Itoa(vigencia) + "."}
	articulos = append(articulos, articulo)
	resolucion.Articulos = articulos //articulos//articulos

	return resolucion
}
