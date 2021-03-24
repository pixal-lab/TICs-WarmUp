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
		/*
			cookie := http.Cookie{Name: "locality", Value: "here", Expires: time.Now().Add(time.Hour), HttpOnly: true, MaxAge: 50000, Path: "/"}
			http.SetCookie(w, &cookie)
			fmt.Fprintf(w, "Yay")
			cooki, err := r.Cookie("locality")
			fmt.Fprint(w, cooki, err) */
		http.Redirect(w, r, "login.html", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
