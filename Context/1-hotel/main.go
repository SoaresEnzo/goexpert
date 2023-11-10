package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	BookHotel(ctx)
}

func BookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout")
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked")
	}
}
