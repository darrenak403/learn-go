package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. Array
	// array.Array()

	// 2. Slice
	//Khoi tao slice
	var numbers []int
	fmt.Println(numbers)

	//[] => Slice rỗng

	//Slice => 0 có kích thước
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//Array => Có kích thước cố định
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)

	fmt.Println("Slice có phải là Slice không ?", reflect.TypeOf(slice).Kind() == reflect.Slice)
	fmt.Println("Array có phải là Array không ?", reflect.TypeOf(array).Kind() == reflect.Array)

	fmt.Println("======Nhung cach khoi tao slice========")
	fmt.Println("======Tao slice tu array========")
	arrV2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arrV2)
	fmt.Println("arrV2 có phải là Array không ?", reflect.TypeOf(arrV2).Kind() == reflect.Array)

	sliceV2 := arrV2[1:4] // 20 30 40
	fmt.Println(sliceV2)
	fmt.Println("sliceV2 có phải là Slice không ?", reflect.TypeOf(sliceV2).Kind() == reflect.Slice)

	fmt.Println("======Tao slice tu make========")
	sliceV3 := make([]int, 2, 5)          // Slice có độ dài 2 và dung lượng 5
	sliceV3[0] = 1                        //[1 0]
	sliceV3[1] = 2                        //[1 2]
	sliceV3 = append(sliceV3, 3, 4, 5, 6) // Thêm phần tử vào slice
	fmt.Println(sliceV3)
	fmt.Println("Chieu dai cua sliceV3", len(sliceV3))
	fmt.Println("Dung luong cua sliceV3", cap(sliceV3))

	fmt.Println("======Duyet slice========")
	apples := []string{"Apple", "Banana", "Cherry"}
	for i, v := range apples {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	school := [][]string{
		{"Math", "Physics", "Chemistry"},
		{"Literature", "History", "Geography"},
	}

	for _, combo := range school {
		for _, subject := range combo {
			fmt.Println(subject)
		}
	}

	fmt.Println("======Hieu ro ve cap va len========")
	sliceV4 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice cha")
	fmt.Println(sliceV4)
	fmt.Println("Chieu dai cua sliceV4", len(sliceV4))
	fmt.Println("Dung luong cua sliceV4", cap(sliceV4))
	fmt.Println("Slice con")
	//subSlice := sliceV4[1:] // 2 3 4 5
	//subSlice := sliceV4[:4] // 1 2 3 4
	subSlice := sliceV4[1:4] // 2 3 4
	fmt.Println(subSlice)
	fmt.Println("Chieu dai cua subSlice", len(subSlice))
	fmt.Println("Dung luong cua subSlice", cap(subSlice))
}
