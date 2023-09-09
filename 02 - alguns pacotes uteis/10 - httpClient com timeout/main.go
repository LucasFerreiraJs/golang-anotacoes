package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "teste"}`))

	resp, err := c.Post("https://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	// faz uma cópia direcionando
	io.CopyBuffer(os.Stdout, resp.Body, nil)

}
