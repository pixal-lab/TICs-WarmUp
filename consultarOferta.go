package main

import (
	"fmt"
	"net/http"
)

func consultarOferta(w http.ResponseWriter, r *http.Request) {
	coockie, err := r.Cookie("user")
	if err != nil {
		fmt.Fprintf(w, "error1")
	}
	container := getOfertas(coockie.Value)
	fmt.Println(container)
}
