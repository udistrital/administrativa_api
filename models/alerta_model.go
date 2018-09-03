package models

type AlertType string

const (
	AlertError   AlertType = "error"
	AlertSucess            = "success"
	AlertWarning           = "warning"
)

type Alert struct {
	Type AlertType
	Code string
	Body interface{}
}
