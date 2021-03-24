package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mux"
	"net/http"
	"time"
)

//json.NewEncoder(w).Encode(" API endpoint  POST")
//recordar que r tiene lso valores provenientes del server y w los valores que se enviaran, no manejo a este momento como trabajar los valores que se envian.
func index(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Print("valor_1 =  " + email + " valor_2 = " + password + "\n")
	if true { //funcion_validadora
		print("si la funcion nos deja pasar entonces ejecutamos AgregarOderta.html pero con los valores de session o algo, hay que hablarlo o simplemente podemos enviar los datos \n ")
		http.Redirect(w, r, "login.html", http.StatusSeeOther)
	} else {
		print("si no nos deja pasar la funcion validadora tendriamos volver a ejecutar el index\n")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

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

/*	Posible forma de registro
	nombre_usuario := r.FormValue("nombre_usuario")
	contraseña := r.FormValue("contraseña")
	validar_contraseña := r.FormValue("validar_contraseña")

	if r.FormValue("contraseña") == r.FormValue("validar_contraseña"){
		http.Redirect(w, r, "index.html", http.StatusSeeOther)
	}else{
		fmt.Printf("No coinciden las contraseñas, vuelva a ingresarlo ")
	} */
func Remover(w http.ResponseWriter, r *http.Request) {
	id_user := r.FormValue("id_user") //buscar id del usuario y el id de la consulta para realizar la  eliminación de esta
	id_consulta := r.FormValue("id_consulta")
	print(id_user, id_consulta)
}
func AgregarOferta(w http.ResponseWriter, r *http.Request) {
	nombre_producto := r.FormValue("nombre_producto")
	monto_solicitado := r.FormValue("monto_solicitado")
	valor_cuotas := r.FormValue("valor_cuotas")
	numero_cuotas := r.FormValue("numero_cuotas")
	/*Da aqui mandar los valores pa la bdd*/
	print("aqui deberian llegar los valores de la oferta \n\n\n")
	print(nombre_producto, monto_solicitado, valor_cuotas, numero_cuotas)
}
func Logout(w http.ResponseWriter, r *http.Request) {

}
func login1(w http.ResponseWriter, r *http.Request) {
	print("paso por el primer boton de login \n ")
	http.Redirect(w, r, "/AgregarOferta.html", http.StatusSeeOther)
}
func login2(w http.ResponseWriter, r *http.Request) {
	print("paso por el primer boton de login \n ")
	http.Redirect(w, r, "/ConsultarOferta.html", http.StatusSeeOther)
}
func consultarOferta(w http.ResponseWriter, r *http.Request) {
	print("prueba\n")
	http.Redirect(w, r, "/ConsultarOferta.html", http.StatusSeeOther)
}

func main() {
	router := mux.NewRouter()
	//los metodos deben llevar la misma palabra en el action del form (html) y en los parametros del HandleFunc
	router.HandleFunc("/", index).Methods("POST")                     //para los datos provenientes del index email y contraseña
	router.HandleFunc("/logout", Logout)                              //Cerrar sesión del usuario
	router.HandleFunc("/login.html1", login1)                         //Cerrar sesión del usuario
	router.HandleFunc("/login.html2", login2)                         //Cerrar sesión del usuario
	router.HandleFunc("/registration.html", registro).Methods("POST") //para los datos provenientes de la pagina de registro email y contraseña
	router.HandleFunc("/AgregarOferta.html", AgregarOferta).Methods("POST")
	router.HandleFunc("/remove", Remover).Methods("POST")
	router.HandleFunc("/ConsultarOferta.html", consultarOferta).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./Maqueta/"))) //ejecuta el /index.html

	server := &http.Server{ // valores predeterminados del server
		Handler:      router,           //variable de mux para el ruteo
		Addr:         ":8080",          // puerto
		WriteTimeout: 15 * time.Second, //tiempo de espera para enviar
		ReadTimeout:  15 * time.Second, //tiempo de espera para leer
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe()) // para errores del servidor, deberia mostrarnos un codigo de estado (ej: 404)
}
