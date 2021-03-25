package main

import (
	"net/http"
)

func Remover(w http.ResponseWriter, r *http.Request) {
	id_user := r.FormValue("id_user") //buscar id del usuario y el id de la consulta para realizar la  eliminación de esta
	id_consulta := r.FormValue("id_consulta")
	userCookie, _ := r.Cookie("user")
	userPw, _ := r.Cookie("pw")
	if id_user == userCookie.Value && loginCheck(id_user, userPw.Value) && ofCheck(id_user, id_consulta) {
		removeOferta(id_consulta)
	}
	http.Redirect(w, r, "./consultar", http.StatusSeeOther)
}
