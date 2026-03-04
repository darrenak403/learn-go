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
func main() {
	myDog, err := dog.New("Bully")
	if err != nil {
		panic(err)
	}
	MakeAnimalSpeak(myDog)

	fmt.Println("-=-=-=-=-=-=-=-=-=--=")

	myCat, err := cat.New("Whiskers")
	if err != nil {
		panic(err)
	}
	MakeAnimalSpeakPlus(myCat)

	fmt.Println("-=-=-=-=-=-=-=-=-=--=")

	myMouse, err := mouse.New("Jerry")
	if err != nil {
		panic(err)
	}

	fmt.Println(myMouse.Run())
}
