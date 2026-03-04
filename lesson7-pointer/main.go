package main

import "fmt"

func lythuyetPointer() {
	name := "Thanh Dat"
	fmt.Println("-=-=-=-=-=-=-= Infomation name variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

	//Tao con tro
	ptrName := &name
	fmt.Println("-=-=-=-=-=-=-= Infomation ptrName variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", ptrName)
	fmt.Printf("Value: %v \n", ptrName)
	fmt.Printf("Address: %v \n", &ptrName)

	fmt.Printf("Find value name of ptrName: %v", *ptrName)

	fmt.Println("-=-=-=-=-=-=-= Infomation ptrName2 variable -=-=-=-=-=")
	//Tao con tro moi
	ptrName2 := &ptrName
	fmt.Printf("Data type: %T \n", ptrName2)
	fmt.Printf("Value: %v \n", ptrName2)
	fmt.Printf("Address: %v \n", &ptrName2)

	fmt.Printf("Find value name of ptrName2: %v", **ptrName2)
}

func updateName(name *string) {
	*name = "Updated Name"
	fmt.Println("-=-=-=-=-=-=-= Infomation name variable in updateName function -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", *name)
	fmt.Printf("Value: %v \n", *name)
	fmt.Printf("Address: %v \n", name)
}

func main() {
	name := "Thanh Dat"
	fmt.Println("-=-=-=-=-=-=-= Infomation name variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

	name = "Thien Kieu"
	fmt.Println("-=-=-=-=-=-=-= Infomation name variable afterUpdate -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

	fmt.Println()

	updateName(&name)

	fmt.Println("-=-=-=-=-=-=-= Infomation name variable after run Update Name func -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

}
