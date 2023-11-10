package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	BookHotel(ctx, "Hotel")
}

func BookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
}
