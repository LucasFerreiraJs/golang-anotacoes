package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	// handleFunc e handle
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	mux.HandleFunc("/teste", HomeHandler)

	// handler
	// sempre implementa uma interface com método serveHTTP
	mux.Handle("/blog", Blog{title: "exemplo handle"})
	http.ListenAndServe(":8080", mux)

	// mux2 := http.NewServeMux()
	// mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("mux 2"))
	// })

	// http.ListenAndServe(":8081", mux2)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type Blog struct {
	title string
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))

}
