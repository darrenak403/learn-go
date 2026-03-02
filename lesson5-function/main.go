package main

import "fmt"

// func sum(numb1 int, numb2 int) {
// 	tong := numb1 + numb2
// 	fmt.Println(tong)
// }

// func sum(numb1, numb2 int) (tong int) {
// 	tong = numb1 + numb2
// 	return
// }

func sum(numb1, numb2 int) (int, int, int, float32) {
	tong := numb1 + numb2
	hieu := numb1 - numb2
	tich := numb1 * numb2
	thuong := float32(numb1) / float32(numb2)
	return tong, hieu, tich, thuong
}

func countDown(number int) {
	if number < 0 {
		fmt.Println("Number must be greater than or equal to 0")
		return
	}
	fmt.Println(number)
	countDown(number - 1)
}

func main() {
	// tong, hieu, tich, thuong := sum(5, 3)
	// fmt.Println(tong, hieu, tich, thuong)
	countDown(10)
}
