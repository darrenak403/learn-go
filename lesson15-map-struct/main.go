package main

import (
	"fmt"
)

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	employees := map[string]Employee{
		"ep1": {
			Name: "Dada",
			Age:  20,
			Role: "Student",
		},
		"ep2": {
			Name: "Dudu",
			Age:  21,
			Role: "Student",
		},
		"ep3": {
			Name: "Didi",
			Age:  22,
			Role: "Student",
		},
	}

	fmt.Printf("%v\n", employees)

	for key, value := range employees {
		fmt.Printf("Gia tri cua key %s  \n", key)
		fmt.Printf("Name: %s \n", value.Name)
		fmt.Printf("Age: %d \n", value.Age)
		fmt.Printf("Role: %s \n", value.Role)
		fmt.Println("--------------------------------")
	}

	studentSubject := map[string][]string{
		"Dada": {"Math", "Physics", "Chemistry"},
		"Dudu": {"Math", "Physics", "Chemistry"},
		"Didi": {"Math", "Physics", "Chemistry"},
	}

	fmt.Printf("%v\n", studentSubject)

	fmt.Printf("Mon hoc cua Dada la: %v \n", studentSubject["Dada"][0])

	for key, value := range studentSubject {
		for _, subject := range value {
			fmt.Printf("Mon hoc cua %s la: %s \n", key, subject)
		}
	}
}
