package main

import (
	"context"
	"time"
)

func employee(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			println("Employee: Received cancellation signal, stopping work.")
			return
		default:
			priority := ctx.Value("priority")
			println("Employee: Working with priority", priority)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "priority", "high")
	go employee(ctx)

	time.Sleep(3 * time.Second)
}
