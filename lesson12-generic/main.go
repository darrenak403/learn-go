package main

import (
	"cmp"
	"fmt"
)

type Box[T any] struct {
	Content     T
	Description T
}

func PrintValue[T any](value T) {
	fmt.Println(value)
}

func IsEqual[T comparable](a, b T) bool {
	return a == b
}

func IsNotEqual[T comparable](a, b T) bool {
	return a != b
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxLengthString(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

func main() {
	//1. Generic function
	// PrintValue(10)
	// PrintValue("Hello, World!")
	// PrintValue(3.14)
	// PrintValue(true)

	//2. Comparable function
	// fmt.Println(IsEqual(5, 5))             // true
	// fmt.Println(IsEqual("hello", "hello")) // true
	// fmt.Println(IsNotEqual(5, 10))         // true

	//3 - Generic function with cmp.Ordered constraint
	// fmt.Println(Max(10, 20))             // 20
	// fmt.Println(Max(3.14, 2.71))         // 3.14s
	// fmt.Println(Max("Dat", "Thanh Dat")) // "Thanh Dat" ASCII value of 'T' is greater than 'D'

	// fmt.Println(MaxLengthString("Dat", "Thanh Dat")) // "Thanh Dat"

	//4. Generic struct
	stringBox := Box[string]{Content: "Hello, World!", Description: "A box that contains a string"}
	intBox := Box[int]{Content: 42, Description: 99}
	PrintValue(stringBox.Content)
	PrintValue(intBox.Content)
	PrintValue(stringBox.Description)
	PrintValue(intBox.Description)
}
