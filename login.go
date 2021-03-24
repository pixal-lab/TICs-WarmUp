package main

import (
	"net/http"
)

//json.NewEncoder(w).Encode(" API endpoint  POST")
//recordar que r tiene lso valores provenientes del server y w los valores que se enviaran, no manejo a este momento como trabajar los valores que se envian.
func login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if loginCheck(email, password) { //funcion_validadora
		cookie := http.Cookie{Name: "user", Value: email}
		http.SetCookie(w, &cookie)
		cookie2 := http.Cookie{Name: "pw", Value: password}
		http.SetCookie(w, &cookie2)

		// rec, _ := r.Cookie("pw")
		// fmt.Fprint(w, rec.Value)
		http.Redirect(w, r, "login.html", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
