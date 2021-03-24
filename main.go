package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")                          //inicia index
	router.HandleFunc("/register", registro).Methods("POST")                    //registration.html
	router.HandleFunc("/addOffert", AgregarOferta).Methods("POST")              //agregaroferta.html
	router.HandleFunc("/ConsultarOferta.html", consultarOferta).Methods("POST") //no existe pero hay que poder eliminar objetos de este
	router.HandleFunc("/logout", Logout)
	router.HandleFunc("/remove", Remover).Methods("POST")

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
