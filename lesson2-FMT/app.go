package main

import "fmt"

func main() {
	// fmt.Print("Hello, World!")
	// fmt.Print("Thanh Dat")

	// fmt.Println("Hello, World!")
	// fmt.Println("Thanh Dat")

	// var fullName string = "Thanh Dat"
	// var age int = 22
	// fmt.Printf("My name is %s, I'm %d years old\n", fullName, age)

	// var fullName, address string
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Please enter your name: ")
	// scanner.Scan()
	// fullName = scanner.Text()
	// fmt.Print("Please enter your address: ")
	// scanner.Scan()
	// address = scanner.Text()
	// fmt.Printf("Your name is %s, you live in %s\n", fullName, address)

	// message := fmt.Sprintln("Hello, ", "Thanh Dat")
	// fmt.Print(message)

	//Bai 5: Kieu du lieu cua bien
	ten := "Thanh Dat"
	tuoi := 22
	chieuccao := 1.75
	daTotNghiep := true
	phanTram := 10

	fmt.Printf("Kieu du lieu cua bien ten: %T\n", ten)
	fmt.Printf("Kieu du lieu cua bien tuoi: %T\n", tuoi)
	fmt.Printf("Kieu du lieu cua bien chieuCao: %T\n", chieuccao)
	fmt.Printf("Kieu du lieu cua bien daTotNghiep: %T\n", daTotNghiep)
	fmt.Printf("Kieu du lieu cua bien phanTram: %T\n", phanTram)

	fmt.Printf("Toi ten la: %v \n", ten)

	fmt.Printf("Toi ten la: %s, toi %d tuoi, chieu cao cua toi la %.2fm, da tot nghiep: %t, phan tram cua toi la: %v\n", ten, tuoi, chieuccao, daTotNghiep, phanTram)

	fmt.Printf("Chieu cao cua toi la %.2fm \n", chieuccao)
	fmt.Printf("Chieu cao cua toi la %.5fm \n", chieuccao)
	fmt.Printf("Chieu cao cua toi la %.1fm \n", chieuccao)

}
