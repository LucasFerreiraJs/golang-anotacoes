package main

import (
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	// irá ser chamado por ultimo
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))
	// req.Body.Close()

}
