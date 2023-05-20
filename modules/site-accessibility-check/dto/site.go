package dto

type Site struct {
	Name     string `json:"name"`
	IsOn     bool   `json:"isOn"`
	TimeinMs int64  `json:"timeInMs"`
}
