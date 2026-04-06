package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() //Giảm giá trị của WaitGroup đi 1
	fmt.Printf("Task %d bat dau \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Task %d ket thuc \n", id)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1) // Tăng giá trị của WaitGroup lên 1
		go task(i, &wg)
	}

	wg.Wait() // Chờ cho đến khi giá trị của WaitGroup bằng 0

	fmt.Printf("Tong thoi gian thuc hien: %s\n", time.Since(start))
}
