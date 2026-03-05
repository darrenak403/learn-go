package main

import (
	"fmt"

	"darrenak.com/learn-go-basics/cat"
	"darrenak.com/learn-go-basics/dog"
	"darrenak.com/learn-go-basics/mouse"
	"darrenak.com/learn-go-basics/service"
)

func MakeAnimalSpeak(a service.Animal) {
	fmt.Printf("Day la ten cua con vat: %s \n", a.GetName())
	fmt.Println(a.Speak())
}

func MakeAnimalSpeakPlus(p service.AnimalPlus) {
	fmt.Printf("Day la ten cua con vat: %s \n", p.GetName())
	fmt.Println(p.Speak())
	fmt.Println(p.Eat())
}

// Empty interface
func PrintValue(val interface{}) {
	// str, ok := val.(string)
	// if ok {
	// 	fmt.Println(str)
	// } else {
	// 	fmt.Println("Please send a string value")
	// }
	switch val := val.(type) {
	case string:
		fmt.Println(val)
	case int:
		fmt.Println(val)
	case float64:
		fmt.Println(val)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	myDog, err := dog.New("Bully")
	if err != nil {
		panic(err)
	}
	MakeAnimalSpeak(myDog)

	PrintValue("-=-=-=-=-=-=-=-=-=--=")

	myCat, err := cat.New("Whiskers")
	if err != nil {
		panic(err)
	}
	MakeAnimalSpeakPlus(myCat)

	PrintValue("-=-=-=-=-=-=-=-=-=--=")

	myMouse, err := mouse.New("Jerry")
	if err != nil {
		panic(err)
	}

	PrintValue(myMouse.Run())

	PrintValue("=--==-=-==-=-==")

	PrintValue(5.8)
	PrintValue(true)
}
