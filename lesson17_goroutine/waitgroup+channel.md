package main

import (
"fmt"
"sync"
"time"
)

func task(id int, ch chan<- string, wg \*sync.WaitGroup) {
defer wg.Done() //Giảm giá trị của WaitGroup đi 1
fmt.Printf("Task %d bat dau \n", id)
time.Sleep(time.Second)
ch <- fmt.Sprintf("Task %d ket thuc ", id)

}

func main() {
start := time.Now()
var wg sync.WaitGroup
ch := make(chan string)

    for i := 1; i <= 4; i++ {
    	wg.Add(1) // Tăng giá trị của WaitGroup lên 1
    	go task(i, ch, &wg)
    }

    go func() {
    	wg.Wait() // Chờ cho đến khi giá trị của WaitGroup bằng 0
    	close(ch)
    }()

    for value := range ch {
    	fmt.Println(value)
    }

    fmt.Printf("Tong thoi gian thuc hien: %s", time.Since(start))

}
