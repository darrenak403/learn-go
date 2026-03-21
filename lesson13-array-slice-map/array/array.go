package array

import (
	"fmt"
)

type Employee struct {
	Id     int
	Name   string
	Age    int
	Salary float64
}

func Array() {

	employees := [...]Employee{
		{Id: 1, Name: "John Doe", Age: 30, Salary: 50000},
		{Id: 2, Name: "Jane Smith", Age: 25, Salary: 60000},
		{Id: 3, Name: "Bob Johnson", Age: 35, Salary: 55000},
	}

	// fmt.Println(employees)
	fmt.Printf("Id: %d \n", employees[1].Id)
	fmt.Printf("Name: %s \n", employees[1].Name)
	fmt.Printf("Age: %d \n", employees[1].Age)
	fmt.Printf("Salary: %.2f \n", employees[1].Salary)
	fmt.Println("===============for-range=================")
	for _, employee := range employees {
		fmt.Printf("Id: %d, Name: %s, Age: %d, Salary: %.2f \n", employee.Id, employee.Name, employee.Age, employee.Salary)
	}

	// var number int
	// fmt.Println(number)

	// var numbers [5]int
	// fmt.Println(numbers)

	// var character string
	// fmt.Println(character)

	// var characters [3]string
	// fmt.Println(characters)

	//Bai 1: Mang mot chieu
	// var numbers [5]int
	// numbers[0] = 8 // 8 0 0 0 0
	// numbers[2] = 10 // 8 0 10 0 0
	// numbers[4] = 15 // 8 0 10 0 15
	// fmt.Println(numbers)

	// var numbers = [5]int{8, 0, 10, 0, 15}
	// fmt.Println(numbers)

	// var numbers = [...]int{8, 0, 10, 0, 15}
	// fmt.Printf("Total array %T \n", numbers)
	// fmt.Println(numbers)

	//Bai 2: Mang da chieu
	// var matrix [2][3]int = [2][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	// fmt.Println(matrix)
	// fmt.Println(matrix[1][1]) // 5

	// matrix[0][2] = 10
	// // 1 2 10
	// // 4 5 6
	// fmt.Println(matrix)

	//Bai 3: for and for range
	//Mot chieu
	// numbers := [5]int{8, 0, 10, 0, 15}
	// // fmt.Println(numbers[0])
	// // fmt.Println(numbers[1])
	// // fmt.Println(numbers[2])
	// // fmt.Println(numbers[3])
	// // fmt.Println(numbers[4])
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// fmt.Println("_+_+_+_+_+_+_+_+_+_+_+_+_+")
	// //Da chieu
	// matrix := [3][4]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// }
	// fmt.Println(len(matrix))
	// for k := 0; k < len(matrix); k++ {
	// 	// fmt.Println(matrix[k])
	// 	for j := 0; j < len(matrix[k]); j++ {
	// 		fmt.Println(matrix[k][j])
	// 	}
	// }
	// fmt.Println("===============for-range=================")
	// numbers2 := [5]int{8, 0, 10, 0, 15}
	// for index, value := range numbers2 {
	// 	fmt.Printf("Index: %d, Value: %d \n", index, value)
	// }

	// matrix2 := [3][4]int{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// }
	// for key, val := range matrix2 {
	// 	for key2, val2 := range val {
	// 		fmt.Printf("Index: [%d][%d], Value: %d \n", key, key2, val2)
	// 	}
	// }

}
