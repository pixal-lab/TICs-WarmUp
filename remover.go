package main

import (
	"net/http"
)

func Remover(w http.ResponseWriter, r *http.Request) {
	id_user := r.FormValue("id_user") //buscar id del usuario y el id de la consulta para realizar la  eliminación de esta
	id_consulta := r.FormValue("id_consulta")
	userCookie, _ := r.Cookie("user")
	userPw, _ := r.Cookie("pw")
	if id_user == userCookie.Value && loginCheck(id_user,userPw.Value) && {
		
	}
	print(id_user, id_consulta)
}
