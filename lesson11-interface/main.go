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
	fmt.Println(val)
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
}
