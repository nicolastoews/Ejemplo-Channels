package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "lenguaje", "Go")
	foo(ctx, "lenguaje")

	ctx = context.WithValue(ctx, "perro", "Gaston")
	foo(ctx, "perro")

	foo(ctx, "color")
}

func foo(ctx context.Context, s string) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("Valor encontrado:", v)
		return
	}
	fmt.Println("Llave no encontrada:", s)
}
