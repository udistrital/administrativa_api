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
	_, err := o.Raw("SELECT * FROM administrativa.resolucion WHERE administrativa.resolucion.id_resolucion=" + idResolucion + ";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	resolucionCompleta := ResolucionCompleta{Id: temp[0].Id, Consideracion: temp[0].ConsideracionResolucion, Preambulo: temp[0].PreambuloResolucion, Vigencia: temp[0].Vigencia, Numero: temp[0].NumeroResolucion, Titulo: temp[0].Titulo}

	var arts []ComponenteResolucion
	_, err2 := o.Raw("SELECT * FROM administrativa.componente_resolucion WHERE resolucion_id=" + idResolucion + " AND tipo_componente like 'Articulo' ORDER BY numero asc;").QueryRows(&arts)
	if err2 == nil {
		fmt.Println("Consulta exitosa")
	}

	var articulos []Articulo

	for _, art := range arts {
		articulo := Articulo{Id: art.Id, Numero: art.Numero, Texto: art.Texto}

		var pars []ComponenteResolucion
		_, err3 := o.Raw("SELECT * FROM administrativa.componente_resolucion WHERE resolucion_id=" + idResolucion + " AND tipo_componente like 'Paragrafo' AND componente_padre=" + strconv.Itoa(articulo.Id) + " ORDER BY numero asc;").QueryRows(&pars)
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

func GetTemplateResolucion(dedicacion, nivel string) (res ResolucionCompleta) {
	var resolucion ResolucionCompleta
	var articulos []Articulo
	var articulo Articulo

	if dedicacion == "HCP" && nivel == "PREGRADO" {

		resolucion = ResolucionCompleta{Consideracion: "Que mediante Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las Universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital.(Artículos 3 y 4).\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes en las modalidades de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional, en virtud del artículo 13 del Acuerdo 011 de Noviembre 15 de 2002 término Fijo por periodos académicos.\n\nQue la presente rige y se aplica únicamente a los docentes de Vinculación Especial Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional en Pregrado en lo pertinente únicamente al valor del punto.”\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de enero 31 de 2003, 0013-A de enero 31 de 2003, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\nQue el artículo 128 de la Carta Política consigna que nadie podrá desempeñar simultáneamente más de un empleo público ni recibir más de una asignación que provenga del tesoro público salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d del artículo 19 de la Ley 4 de 1992 determinó como excepción los honorarios percibidos por concepto hora cátedra.\n\nQue mediante Acuerdo 006 de julio 19 de 2002 se fija el valor de la Hora Cátedra por Honorarios y se establece un número máximo de horas para los docentes de carrera que presten servicios a la Universidad Distrital Francisco José de Caldas, en los programas de Postgrado, y modifican parcialmente los Acuerdos 005 y 007 de 2001.\n\nQue mediante Acuerdo 002 de enero 31 de 2003 se modifica y reglamenta el Acuerdo 001 de enero 17 de 2003.\n\nQue para efectos de pago de Salarios y prestaciones sociales, el periodo académico corresponde de conformidad con la Resolución No 182 de Diciembre 19 de 2017 emitida por el Consejo Académico, por la cual se expide para los proyectos curriculares de pregrado y posgrado de la Universidad Distrital Francisco José de Caldas el calendario académico para el año 2018 comprendido entre el 1° de Agosto de 2018 al 11 de Diciembre del 2018 equivalentes a 4,5 meses (mes de cuatro semanas) y el mes comprenderá (30) días.\n\nQue mediante Resolución No. 001 del 15 de Febrero de 2012 emitida por la vicerrectoría Académica se establecen los procesos de selección y vinculación de docentes de vinculación especial.\n\nQue de conformidad con el artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018 establece “A partir del 1° de enero de 2018, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en trece mil quinientos noventa y ocho pesos ($13.598) moneda corriente”.\n\nQue mediante Resolución 072 del quince de marzo del 2018 emitida por Rectoría mediante la cual acoge únicamente el valor del punto para los docentes de vinculación especial a quienes se les vinculó o reconoció honorarios en esta vigencia con el valor del punto anterior, y se hace necesario dar cumplimiento al artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018.\n\nQue para efectos presupuestales de la presente resolución se hará a cargo de la Disponibilidad presupuestal No XXXX del XX de XXXXX de 2018.\n\nEn virtud de lo anteriormente expuesto,\n\n"}

		articulo = Articulo{Texto: "Vincular para el Tercer Periodo Académico del año 2018 como docentes en la modalidad de Hora Cátedra de Vinculación Especial en el escalafón y dedicación establecidas en la tabla, a los siguientes docentes:	"}
		articulos = append(articulos, articulo)
	}

	if dedicacion == "HCH" && nivel == "PREGRADO" {

		resolucion = ResolucionCompleta{Consideracion: "Que mediante Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las Universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital.(Artículos 3 y 4).\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes en las modalidades de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional, en virtud del artículo 13 del Acuerdo 011 de Noviembre 15 de 2002 término Fijo por periodos académicos.\n\nQue la presente rige y se aplica únicamente a los docentes de Vinculación Especial Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional en Pregrado en lo pertinente únicamente al valor del punto.”\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de enero 31 de 2003, 0013-A de enero 31 de 2003, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\nQue de conformidad con el artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018 establece “A partir del 1° de enero de 2018, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en trece mil quinientos noventa y ocho pesos ($13.598) moneda corriente”.\n\nQue mediante Resolución 001 del 15 de febrero del 2012 emitida por la Vicerrectoría Académica establece el proceso de selección y vinculación de Docentes en la modalidad de Hora Cátedra, Honorarios, Medio Tiempo Ocasional y Tiempo Completo Ocasional.\n\nQue de acuerdo a lo establecido en el artículo 5º del Decreto 318 de diecinueve (19) de febrero de 2018, los Rectores Universitarios expedirán los correspondientes actos administrativos.\n\nQue para efectos de pago de Honorarios, el periodo académico  corresponde de conformidad con la Resolución No 182 de Diciembre 19 de 2017 emitida por el Consejo Académico, por la cual se expide para los proyectos curriculares de pregrado y posgrado de la Universidad Distrital Francisco José de Caldas el calendario académico para el año 2018 comprendido entre el 1° de Agosto de 2018 al 11 de Diciembre del 2018 equivalentes a 4,5 meses (mes de cuatro semanas) y el mes comprenderá (30) días.\n\nQue mediante Resolución 072 del quince de marzo del 2018 emitida por Rectoría mediante la cual acoge únicamente el valor del punto para los docentes de vinculación especial a quienes se les vinculó o reconoció honorarios en esta vigencia con el valor del punto anterior, y se hace necesario dar cumplimiento al artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018.\n\nQue para efectos presupuestales de la presente resolución se hará a cargo de la disponibilidad presupuestal No xxxx del xx de xxxxxxxxxx de 2018.\n\nEn virtud de lo anteriormente expuesto,\n\n"}

		articulo = Articulo{Texto: "Reconocer Honorarios para el Tercer Periodo Académico del año 2018, como docentes de Hora Cátedra Honorarios (Vinculación Especial) en el escalafón y dedicación establecidas en la tabla, a los siguientes docentes:	"}
		articulos = append(articulos, articulo)
	}

	if dedicacion == "HCP" && nivel == "POSGRADO" {

		resolucion = ResolucionCompleta{Consideracion: "Que mediante Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las Universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital. (Artículos 3 y 4).\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes en las modalidades de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional, en virtud del artículo 13 del Acuerdo 011 de noviembre 15 de 2002. Término fijo por periodos académicos.\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones: 317 de septiembre 8 de 2006 y 318 de Septiembre 8 de 2006, Acuerdos 005 y 007 de 2001, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y Ley 4 de 1992).\n\nQue el  artículo 128 de la Carta Política consigna que nadie podrá desempeñar simultáneamente más de un empleo público ni recibir más de una asignación que provenga del tesoro público salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d del artículo 19 de la Ley 4 de 1992 determinó como excepción los honorarios percibidos por concepto hora cátedra.\n\nQue mediante Acuerdo 006 de julio 19 de 2002 se fija el valor de la Hora Cátedra por Honorarios y se establece un número máximo de horas para los docentes de carrera que presten servicios a la Universidad Distrital Francisco José de Caldas, en los programas de Postgrado, y modifican parcialmente los Acuerdos 005 y 007 de 2001.\n\nQue mediante Acuerdo 002 de enero 31 de 2003 se modifica y reglamenta el Acuerdo 001 de enero 17 de 2003.\n\nQue mediante Acuerdo 005 de julio 27 de 2001 se establece que por medio del cual se fija la Hora Cátedra y se establece el número máximo de horas, para los docentes que presten servicios a la Universidad Distrital en los programas de Postgrado.\n\nQue para efectos de pago, reconocimiento y liquidación, el periodo académico corresponde de conformidad con el calendario académico emitido por el Consejo Académico mediante Resolución 182 de Diciembre 19 de 2017 para las actividades del año 2018 a desarrollarse en los programas de posgrado en la Universidad, para este caso se reconocerán de conformidad con los módulos de las horas efectivamente dictadas y aprobadas en los diferentes programas como carga lectiva.\n\nQue mediante Resolución No. 001 del 15 de Febrero de 2012 emitida por la vicerrectoría Académica se establecen los procesos de selección y vinculación de docentes de vinculación especial.\n\nQue para efectos presupuestales de la presente resolución, se hará a cargo de la disponibilidad Presupuestal No. XXXX del XX de XXXX de 2018.\n\nEn virtud de lo anteriormente expuesto,\n\n"}

		articulo = Articulo{Texto: "Vincular para el TERCER PERIODO ACADEMICO del 2018, docentes de Vinculación Especial en el escalafón y dedicación establecidas en la tabla, a los siguientes docentes:	"}
		articulos = append(articulos, articulo)
	}

	if dedicacion == "HCH" && nivel == "POSGRADO" {

		resolucion = ResolucionCompleta{Consideracion: "Que mediante Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las Universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital.(Artículos 3 y 4).\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes en las modalidades de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional, en virtud del artículo 13 del Acuerdo 011 de noviembre 15 de 2002 término Fijo por periodos académicos.\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de enero 31 de 2003, 0013-A de enero 31 de 2003, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\nQue el  artículo 128 de la Carta Política consigna que nadie podrá desempeñar simultáneamente más de un empleo público ni recibir más de una asignación que provenga del tesoro público salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d del artículo 19 de la Ley 4 de 1992 determinó como excepción los honorarios percibidos por concepto hora cátedra.\n\nQue mediante Acuerdo 006 de julio 19 de 2002 se fija el valor de la Hora Cátedra por Honorarios y se establece un número máximo de horas para los docentes de carrera que presten servicios a la Universidad Distrital Francisco José de Caldas, en los programas de Postgrado, y modifican parcialmente los Acuerdos 005 y 007 de 2001.\n\nQue mediante Acuerdo 005 de julio 27 de 2001 se establece que por medio del cual se fija la Hora Cátedra y se establece el número máximo de horas, para los docentes que presten servicios a la Universidad Distrital en los programas de Postgrado.\n\nQue mediante Acuerdo 002 de Marzo 17 de 2011 se modifica el artículo 43 del Acuerdo 011 de 2002, por el cual se pueden vincular a personas pensionadas en la modalidad de docentes de hora- cátedra por honorarios.\n\nQue para efectos de pago, reconocimiento y liquidación, el periodo académico corresponde de conformidad con el calendario académico emitido por el Consejo Académico mediante Resolución 182 de Diciembre 19 de 2017 para las actividades del año 2018, a desarrollarse en los programas de posgrado en la Universidad, que para este caso se reconocerán de conformidad con los módulos de las horas efectivamente dictadas y aprobadas en los diferentes programas como carga lectiva.\n\nQue mediante Resolución No. 001 del 15 de Febrero de 2012  emitida por la vicerrectoría Académica se establecen los procesos de selección y vinculación de docentes de vinculación especial.\n\nQue para efectos presupuestales de la presente resolución se hará a cargo de la disponibilidad Presupuestal No XX del XX de XXXXXXX de 2018.\n\nEn virtud de lo anteriormente expuesto,\n\n"}

		articulo = Articulo{Texto: "Reconocer Honorarios para el Tercer Periodo Académico de 2018, como docente de Hora Cátedra por Honorarios en el escalafón y dedicación establecidos en la tabla, a los siguientes docentes:"}
		articulos = append(articulos, articulo)
	}

	if dedicacion == "TCO-MTO" {

		resolucion = ResolucionCompleta{Consideracion: "Que mediante Decreto 1279 del 19 de Junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las Universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital.\n\nQue la Universidad Distrital Francisco José de Caldas puede vincular docentes en las modalidades de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional, en virtud del artículo 13 del Acuerdo 011 de Noviembre 15 de 2002. Término Fijo por periodos académicos.\n\nQue la presente rige y se aplica únicamente a los docentes de Vinculación Especial Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional en Pregrado en lo pertinente únicamente al valor del punto.”\n\nQue los servicios de los Docentes de Vinculación Especial señalados en el Acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 318 de septiembre de 2006, 317 de septiembre 8 de 2006, Ley 30 de 1992 y Acuerdo 003 de 1997, Artículo 49 y ley 4 de 1992).\n\nQue el  artículo 128 de la Carta Política consigna que nadie podrá desempeñar simultáneamente más de un empleo público ni recibir más de una asignación que provenga del tesoro público salvo las excepciones establecidas en la Ley.\n\nQue en virtud de la anterior norma constitucional, el literal d) del artículo 19 de la Ley 4 de 1992 determinó como excepción los honorarios percibidos por concepto hora cátedra.\n\nQue mediante Acuerdo 002 de enero 31 de 2003 se modifica y reglamenta el Acuerdo 001 de enero 17 de 2003.\n\nQue para efectos de pago de Salarios y prestaciones sociales, el periodo académico corresponde de conformidad con la Resolución No 182 de Diciembre 19 de 2017 emitida por el Consejo Académico, por la cual se expide para los proyectos curriculares de pregrado y posgrado de la Universidad Distrital Francisco José de Caldas el calendario académico para el año 2018 comprendido entre el 1° de Agosto 2018 al 11 de Diciembre del 2018 equivalentes a 4,5 meses (mes de cuatro semanas) y el mes comprenderá (30) días.\n\nQue de conformidad con el artículo 2º del Decreto 318 de diecinueve (19) de febrero de 2018 establece “A partir del 1° de enero de 2018, fijase el valor del punto para los empleados públicos docentes a quienes se les aplica el decreto 1279 del 2002, y demás disposiciones que lo modifiquen o adicionen en trece mil quinientos noventa y ocho pesos ($13.598) moneda corriente”.\n\nQue mediante Resolución 075 del quince de marzo del 2018 emitida por Rectoría mediante la cual acoge únicamente el valor del punto para los docentes de vinculación especial a quienes se les vinculó o reconoció honorarios en esta vigencia con el valor del punto anterior, y se hace necesario dar cumplimiento al artículo 2° del Decreto 318 de diecinueve (19) de febrero de 2018.\n\nQue mediante Resolución 001 del 15 de Febrero del 2012 emitido por la Vicerrectoría Académica establece el proceso de selección y vinculación de Docentes en la modalidad de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Completo Ocasional.\n\nQue para efectos presupuestales de la presente resolución se hará a cargo de la disponibilidad presupuestal No XXXX del XX de XXXXX de 2018.\n\nEn virtud de lo anteriormente expuesto,\n\n"}
		articulo = Articulo{Texto: "Vincular para el Tercer Periodo Académico del año 2018, como docentes de Medio Tiempo Ocasional y Tiempo Completo Ocasional (Vinculación Especial) en el escalafón y dedicación establecidas en la tabla, a los siguientes docentes: "}
		articulos = append(articulos, articulo)
	}

	articulo = Articulo{Texto: "El pago de los servicios prestados por los profesores de Vinculación Especial según su escalafón, se cancelarán previa certificación de las horas efectivamente dictadas, expedida por el Decano(a) y/o Director(a) o quien haga las veces."}
	paragrafo := Paragrafo{Texto: "El valor del punto en pesos para el reconocimiento y pago de los docentes de Hora Cátedra, será el que fije el Gobierno Nacional, mediante Decreto cada año y que la Universidad acoja mediante acto administrativo para los docentes de Vinculación Especial."}
	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "El docente deberá cumplir con las obligaciones inherentes a la naturaleza del servicio, contempladas en la Ley, en los Reglamentos de la Universidad y en los Planes de Trabajo entregados por el Profesor y aprobados por el Decano y/o Director"}
	if dedicacion == "HCH" {
		paragrafo = Paragrafo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad mediante acto administrativo hará la liquidación con corte a la fecha del cumplido expedido por el Decano."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
		paragrafo = Paragrafo{Texto: "El docente cumplirá con lo establecido en la Circular No. 004 de julio 19 de 2017 de Rectoría con relación al pago de aportes a seguridad social de forma independiente de conformidad con la ley."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	} else {
		paragrafo = Paragrafo{Texto: "En caso de incumplimiento o retiro del docente, la Universidad mediante acto administrativo hará la liquidación con corte a la fecha del cumplido expedido por el Decano y se cancelarán las prestaciones sociales en la última liquidación del periodo académico que efectúe la División de Recursos Humanos."}
		articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	}
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "El gasto que ocasione la presente resolución se hará con cargo al presupuesto de la actual vigencia, previa certificación de disponibilidad presupuestal."}
	paragrafo = Paragrafo{Texto: "En todo caso, los pagos correspondientes estarán sujetos a las apropiaciones presupuestales y a las transferencias de la Secretaría de Hacienda Distrital."}
	articulo.Paragrafos = append(articulo.Paragrafos, paragrafo)
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "Comunicar de la presente resolución al Docente, quien manifestará bajo la gravedad de juramento que no se encuentra incurso en inhabilidades e incompatibilidades de ley para aceptar dicha vinculación especial, que no tienen con cruces de horarios ni ostentan otra vinculación de carácter público, diferente a hora cátedra en entidades de educación oficiales, siempre y cuando los honorarios sumados no correspondan a más de ocho (8) horas diarias de trabajo, que no ostente vinculación de Tiempo Completo en dos entidades ni Medio Tiempo Ocasional únicamente a excepción de hora cátedra."}
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "En caso de declaratoria de suspensión de actividades académicas, por parte de los máximos organismos de gobierno de la Universidad y autorización previa del Ministerio de Trabajo, cesará para el docente la vinculación especial la obligación de prestar el servicio y para la Universidad la de pagar los salarios, pero persistirá para esta última, la de continuar efectuando los respectivos aportes a salud y pensión en el porcentaje que le corresponda."}
	articulos = append(articulos, articulo)
	articulo = Articulo{Texto: "La presente Resolución se expide a los XXXXXXX (XX) días del mes de XXXXXX de 2018 y surte efectos para el Tercer Periodo Académico del año 2018."}
	articulos = append(articulos, articulo)
	resolucion.Articulos = articulos //articulos//articulos

	return resolucion
}
