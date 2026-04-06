package main

import (
	"context"
	"fmt"
	"time"
)

func cookPho(ctx context.Context, chPho chan<- string) {
	fmt.Println("Bắt đầu nấu phở...")
	select {
	case <-time.After(1 * time.Second):
		chPho <- "Phở đã nấu xong!"
	case <-ctx.Done():
		fmt.Println("Nấu phở bị hủy bỏ!")
		return
	}
}

func cookPizza(ctx context.Context, chPho chan<- string) {
	fmt.Println("Bắt đầu nấu pizza...")
	select {
	case <-time.After(2 * time.Second):
		chPho <- "Pizza đã nấu xong!"
	case <-ctx.Done():
		fmt.Println("Nấu pizza bị hủy bỏ!")
		return
	}
}

func main() {
	chPho := make(chan string)
	chPizza := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	go cookPho(ctx, chPho)
	go cookPizza(ctx, chPizza)

	for i := 0; i < 2; i++ {
		select {
		case pho := <-chPho:
			fmt.Println("Nhan duoc:", pho)
		case pizza := <-chPizza:
			fmt.Println("Nhan duoc:", pizza)
		case <-ctx.Done():
			fmt.Println("Hết thời gian chờ, hủy bỏ nấu ăn!")
			return
		}
	}
}
