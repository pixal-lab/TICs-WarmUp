package main

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
