package main

import (
"fmt"
)

func main() {
ch := make(chan int, 3)

    ch <- 10 //Block -> Gửi
    ch <- 20 //Block -> Gửi
    ch <- 30 //Block -> Gửi

    //Goroutine Anonymous
    go func() {
    	defer close(ch)
    	ch <- 10 //Block -> Gửi
    	ch <- 20 //Block -> Gửi
    	ch <- 30 //Block -> Gửi
    }()

    //Nhận dữ liệu từ channel
    //Chỉ dùng được cách này nếu biết chắc chắn channel sẽ được đóng
    for value := range ch {
    	fmt.Println(value)
    }

    //Dùng cách này nếu biết chắc chắn số lần gửi dữ liệu
    // for i := 1; i <= 3; i++ {
    // 	fmt.Println(<-ch)
    // }

    // time.Sleep(1 * time.Second)

}
