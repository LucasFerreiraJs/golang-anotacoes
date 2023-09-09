package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha123")
	handleHotel(ctx, "Hotel")

}

// varáveis sobre contexto sempre vem primeiro
func handleHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
	fmt.Println(name)
}
