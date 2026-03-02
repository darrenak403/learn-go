package main

import "fmt"

func main() {
	s1 := 15
	s2 := 4

	tong := s1 + s2
	hieu := s1 - s2
	tich := s1 * s2
	thuong := float32(s1) / float32(s2)
	mod := s1 % s2

	fmt.Printf("Tong cua %d va %d la %d \n", s1, s2, tong)
	fmt.Printf("Hieu cua %d va %d la %d \n", s1, s2, hieu)
	fmt.Printf("Tich cua %d va %d la %d \n", s1, s2, tich)
	fmt.Printf("Thuong cua %d va %d la %.2f \n", s1, s2, float64(thuong))
	fmt.Printf("Mod cua %d va %d la %d \n", s1, s2, mod)

	// Bai 2: Toan tu gan
	a := 10
	a += 5 // a = a + 5
	fmt.Printf("Gia tri cua a sau khi su dung toan tu gan: %d \n", a)

	// Bai 3: Toan tu tang giam
	b := 10
	b++ // b = b + 1
	fmt.Printf("Gia tri cua b sau khi su dung toan tu tang: %d \n", b)
	c := 10
	c-- // c = c - 1
	fmt.Printf("Gia tri cua c sau khi su dung toan tu giam: %d \n", c)

	// Bai 4: Toan tu so sanh
	d1 := 10
	d2 := 20
	fmt.Printf("d1 == d2: %t \n", d1 == d2)
	fmt.Printf("d1 != d2: %t \n", d1 != d2)
	fmt.Printf("d1 > d2: %t \n", d1 > d2)
	fmt.Printf("d1 < d2: %t \n", d1 < d2)
	fmt.Printf("d1 >= d2: %t \n", d1 >= d2)
	fmt.Printf("d1 <= d2: %t \n", d1 <= d2)

}
