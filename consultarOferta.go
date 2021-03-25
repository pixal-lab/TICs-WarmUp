package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type datos struct {
	Vendedor string
	Total    string
	Periodo  string
	Cuota    string
	Cae      string
	User     string
	Id       string
}

type tabla struct {
	Arr      []datos
	Producto string
}
type dat struct {
	Ar []tabla
}

func consultarOferta(w http.ResponseWriter, r *http.Request) {
	var arr1 dat
	userCookie, _ := r.Cookie("user")
	coockie, err := r.Cookie("user")
	if err != nil {
		//log.Println(err)
		fmt.Fprintf(w, "error1")
	}
	container := getOfertas(coockie.Value)

	for _, y := range container {
		var parseId primitive.ObjectID = y["_id"].(primitive.ObjectID)
		var id string = parseId.Hex()
		persona := datos{
			y["vendedor"].(string),
			strconv.Itoa(int(y["total"].(int32))),
			strconv.Itoa(int(y["periodo"].(int32))),
			strconv.Itoa(int(y["cuota"].(int32))),
			fmt.Sprint(y["cae"].(float64)),
			userCookie.Value,
			id}
		prod := y["producto"].(string)
		seguir := true
		for j, t := range arr1.Ar {
			if t.Producto == prod {
				t.Arr = append(t.Arr, persona)
				arr1.Ar[j].Arr = t.Arr
				seguir = false
				break
			}
		}
		if seguir {
			var List []datos
			List = append(List, persona)
			tab := tabla{List, prod}
			arr1.Ar = append(arr1.Ar, tab)
		}
	}
	for _, t := range arr1.Ar {
		sort.Slice(t.Arr[:], func(i, j int) bool {
			return t.Arr[i].Cae < t.Arr[j].Cae
		})
	}
	tmpl, err := template.ParseFiles("./Maqueta/ConsultarOferta.html")
	tmpl.Execute(w, arr1)
}
