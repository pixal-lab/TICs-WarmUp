package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func AgregarOferta(w http.ResponseWriter, r *http.Request) {
	producto := r.FormValue("producto")
	vendedor := r.FormValue("vendedor")
	total, err1 := strconv.Atoi(r.FormValue("total"))
	cuota, err2 := strconv.Atoi(r.FormValue("cuota"))
	periodo, err3 := strconv.Atoi(r.FormValue("periodo"))
	if err1 != nil && err2 != nil && err3 != nil {
		fmt.Fprintf(w, "error")
	} else {
		if validarDatosCalculo(total, cuota, periodo) {
			addOferta("editarConCookies", producto, vendedor, total, cuota, periodo, cae(total, cuota, periodo))
		} else {
			fmt.Fprintf(w, "error2")
		}
	}
}
