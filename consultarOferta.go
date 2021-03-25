package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"text/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type datos struct {
	cuota, total, periodo  int32
	cae                    float64
	vendedor, producto, id string
}

type tabla struct {
	producto string
	arr      []datos
}

var arr1 []tabla

func consultarOferta(w http.ResponseWriter, r *http.Request) {
	coockie, err := r.Cookie("user")
	if err != nil {
		//log.Println(err)
		fmt.Fprintf(w, "error1")
	}
	container := getOfertas(coockie.Value)
	log.Println(container)
	for x, y := range container {
		fmt.Println(x)
		fmt.Println(y["cuota"].(int32), y["vendedor"].(string), y["cae"].(float64)*100)
		var parseId primitive.ObjectID = y["_id"].(primitive.ObjectID)
		var id string = parseId.Hex()
		persona := datos{y["cuota"].(int32), y["total"].(int32), y["periodo"].(int32), y["cae"].(float64), y["vendedor"].(string), y["producto"].(string), id}
		prod := y["producto"].(string)
		seguir := true
		for _, t := range arr1 {
			if t.producto == prod {
				t.arr = append(t.arr, persona)
				seguir = false
				break
			}
		}
		if seguir {
			var List []datos
			List = append(List, persona)
			tab := tabla{prod, List}
			arr1 = append(arr1, tab)
		}
	}
	for _, t := range arr1 {
		sort.Slice(t.arr[:], func(i, j int) bool {
			return t.arr[i].cae < t.arr[j].cae
		})
	}
	tmpl, _ := template.ParseFiles("./AgregarOferta.html")
	tmpl.Execute(w, arr1)
}
