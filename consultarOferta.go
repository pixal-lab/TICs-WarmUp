package main

import (
	"html/template"
	"net/http"
	"os"
	"reflect"
)

var templatefuncs = template.FuncMap{"rangeStruct": RangeStructer}

var htmlTemplate = `{{range .}}<tr>
{{range rangeStruct .}} <td>{{.}}</td>
{{end}}</tr>
{{end}}`

func consultarOferta(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	print("prueba\n")
	http.Redirect(w, r, "/ConsultarOferta.html", http.StatusSeeOther)
	container := getOfertas(user)
	t := template.New("t").Funcs(templatefuncs)
	t, err := t.Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, container)
	if err != nil {
		panic(err)
	}
}
func RangeStructer(args ...interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}
	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Struct {
		return nil
	}
	out := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		out[i] = v.Field(i).Interface()
	}
	return out
}
