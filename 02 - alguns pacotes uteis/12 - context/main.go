package main

import (
	"context"
	"time"
)

func main() {

	// qualquer coisa que utiliza esse contexto que ultrapasse 1s será cancelado
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

}
