package models

import "github.com/manifoldco/promptui"

func GetDefaultMenuTemplate() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "{{ .Label | green }}",
		Inactive: "{{ .Label }}",
		Selected: "{{ .Label }}",
		Details: `{{if .Description}}
--------- Detail ----------

{{.Description}}{{end}}`,
	}
}
