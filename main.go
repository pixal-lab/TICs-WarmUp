package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func usuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>por aqui me mato</h1>")
}

func loco(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>como tamo</h1>")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {

	msg := mensaje{
		msg: "porfavor matenme",
	}
	r := http.NewServeMux()
	nuevo := http.FileServer(http.Dir("Maqueta"))
	//---------Funciones de prueba
	r.Handle("/", nuevo)
	r.HandleFunc("/usuario", usuario)
	r.HandleFunc("/loco", loco)
	r.Handle("/hola", msg)

	server := &http.Server{
		Addr:           ":8080",          //puerto de entrada
		Handler:        r,                //la función que ayudará a iniciaar
		ReadTimeout:    10 * time.Second, //tiempo de lectura (10 segundos ambos)
		WriteTimeout:   10 * time.Second, //tiempo de respuesta (si no se declara afecta el rendimiento)
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()
} //
