package templates

import "html/template"

var Cache *template.Template

//go:generate go-bindata -prefix "views/" -pkg templates -o bindata.go views/...

func init(){
	Cache=template.Must(template.ParseFiles(
		"views/form.html",
		"views/error.html",
))
}
