package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		// imprime no command line
		log.Println("Request processada com sucesso")
		// imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		// imprime no command line
		log.Println("Request cancelada pelo cliente")

	}
}
