package main

import (
	"net/http"
	"time"
)

func logout(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "user", Value: "0", Expires: time.Now().Add(-time.Duration(1) * time.Minute), MaxAge: -1}
	http.SetCookie(w, &c)
	c2 := http.Cookie{Name: "pw", Value: "0", Expires: time.Now().Add(-time.Duration(1) * time.Minute), MaxAge: -1}
	http.SetCookie(w, &c2)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
