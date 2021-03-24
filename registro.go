package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func registro(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Print("valor_1 =  " + email + " valor_2 = " + password + "\n")
	if true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		json.NewEncoder(w).Encode("usuario no disponible")
	}
}
