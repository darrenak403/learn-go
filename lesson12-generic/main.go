package main

import "fmt"

func PrintValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	PrintValue(10)
	PrintValue("Hello, World!")
	PrintValue(3.14)
	PrintValue(true)
}
