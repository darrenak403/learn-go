package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done() // Báo cho WaitGroup rằng công việc đã hoàn thành
	sum := 0
	for i := 0; i < 100e8; i++ {
		sum += i
	}
}

func main() {
	numCPU := runtime.NumCPU() // Lấy số lượng CPU có sẵn
	fmt.Printf("So luong CPU: %d\n", numCPU)

	runtime.GOMAXPROCS(numCPU) // Thiết lập số lượng CPU mà Go có thể sử dụng

	start := time.Now() // Lấy thời gian bắt đầu

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go heavyTask(&wg) // Chạy công việc nặng trong một goroutine
	}

	wg.Wait() // Chờ cho tất cả các goroutine hoàn thành

	fmt.Println("Tong thoi gian", time.Since(start))
}
