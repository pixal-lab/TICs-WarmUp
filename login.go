package main

import (
	"fmt"
	"net/http"
)

//json.NewEncoder(w).Encode(" API endpoint  POST")
//recordar que r tiene lso valores provenientes del server y w los valores que se enviaran, no manejo a este momento como trabajar los valores que se envian.
func login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Print("valor_1 =  " + email + " valor_2 = " + password + "\n")
	if loginCheck(email, password) { //funcion_validadora
		http.Redirect(w, r, "login.html", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
