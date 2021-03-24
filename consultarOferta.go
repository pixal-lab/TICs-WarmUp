package main

import (
	"net/http"
)

func consultarOferta(w http.ResponseWriter, r *http.Request) {
	print("prueba\n")
	http.Redirect(w, r, "/ConsultarOferta.html", http.StatusSeeOther)
}
