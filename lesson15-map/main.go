package main

import "fmt"

func main() {
	drink := map[string]int{
		"coke":  10000,
		"pepsi": 12000,
		"sting": 15000,
	}

	fmt.Printf("%v \n", drink)

	student := map[int]string{
		1: "Dada",
		2: "Dudu",
		3: "Didi",
	}

	fmt.Printf("%v \n", student)

	m := make(map[string]int)

	m["coke"] = 10000
	m["pepsi"] = 12000
	m["sting"] = 15000

	fmt.Printf("%v \n", m)

	var food map[string]int

	fmt.Printf("%v \n", food)

	food = make(map[string]int)

	food["coke"] = 10000
	food["pepsi"] = 12000
	food["sting"] = 15000

	fmt.Printf("%v \n", food)

	value, exsits := food["chao"]
	if exsits {
		fmt.Println(value)
	} else {
		fmt.Println("Khong cos trong map")
	}

	for key, value := range food {
		fmt.Printf("%s: %d\n", key, value)
	}
}
