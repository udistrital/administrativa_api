package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

func tiempo_bogota() time.Time {
	fmt.Println("tiempo antes de correccion")
	var tiempoBogota = time.Now()
	logs.Info(tiempoBogota)

	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)
	tiempoBogota = tiempoBogota.In(loc)
	fmt.Println("tiempo despues de correccion")
	logs.Info(tiempoBogota)
	tiempoBogota = tiempoBogota.Add(-10 * time.Hour)
	return tiempoBogota
}

func GetUtcBogota() *time.Location {
	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		fmt.Println(err)
	}
	return loc
}
