package main

import (
	"net/http"
)

func AgregarOferta(w http.ResponseWriter, r *http.Request) {
	nombre_producto := r.FormValue("nombre")
	monto_solicitado := r.FormValue("oferta")
	valor_cuotas := r.FormValue("numero1")
	cuotas := r.FormValue("cuota")
	numero_cuotas := r.FormValue("ncuota")
	/*Da aqui mandar los valores pa la bdd*/
	print("aqui deberian llegar los valores de la oferta \n\n\n")
	print(nombre_producto, monto_solicitado, valor_cuotas, cuotas, numero_cuotas)
}
