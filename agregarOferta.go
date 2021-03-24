package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func AgregarOferta(w http.ResponseWriter, r *http.Request) {
	userCookie, _ := r.Cookie("user")
	producto := r.FormValue("producto")
	vendedor := r.FormValue("vendedor")
	total, err1 := strconv.Atoi(r.FormValue("total"))
	cuota, err2 := strconv.Atoi(r.FormValue("cuota"))
	periodo, err3 := strconv.Atoi(r.FormValue("periodo"))
	if err1 != nil && err2 != nil && err3 != nil {
		fmt.Fprintf(w, "error")
	} else {
		if validarDatosCalculo(total, cuota, periodo) {
			addOferta(userCookie.Value, producto, vendedor, total, cuota, periodo, cae(total, cuota, periodo))
			http.Redirect(w, r, "login.html", http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "error2")
		}
	}
}
