package main

import "fmt"

var (
	name string
	info string
)

func main() {
	//Bài 1:
	// fmt.Println("Hello, Thanh Dat")
	// randomUser()

	//Bài 2:
	var fullName string = "Thanh Dat"
	fmt.Println("Hello, " + fullName)

	var age int
	age = 20
	fmt.Println("I am " + fmt.Sprint(age) + " years old")

	//short variable declaration
	phone := "0123456789"
	fmt.Println("My phone number is " + phone)

	// var math, physics, chemistry float64
	// math = 8.5
	// physics = 7.0
	// chemistry = 9.0
	// average := (math + physics + chemistry) / 3
	// fmt.Printf("The average score is %.2f\n", average)

	math, physics, chemistry := 9.0, 8.5, 7.5
	average := (math + physics + chemistry) / 3
	fmt.Printf("The average score is %.2f\n", average)

	name = "Laptrinh"
	info = "Laptrinh is a programming language."
	fmt.Println(name)
	fmt.Println(info)
}
