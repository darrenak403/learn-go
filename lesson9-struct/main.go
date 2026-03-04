package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Giangvien struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Phone  string `json:"phone"`
}

// khi sử dụng con trỏ struct thì có thể sử dụng trực tiếp tên biến struct để truy cập vào các trường của struct mà không cần phải dereference con trỏ đó
func (gv *Giangvien) showInfoGiangvien() {
	fmt.Printf("Ho ten: %s \n", gv.Name)
	fmt.Printf("Gioi tinh: %d \n", gv.Gender)
	fmt.Printf("So dien thoi: %s \n", gv.Phone)
}

func (gv *Giangvien) clear() {
	gv.Name = ""
	gv.Gender = 0
	gv.Phone = ""
}

func main() {
	thanhdat := Giangvien{
		Name:   "Thanh Dat",
		Gender: 1,
		Phone:  "0123456789",
	}

	// thanhdat.showInfoGiangvien()

	// fmt.Println()

	// thanhdat.clear()

	// fmt.Println("After clear info thanhdat: ")
	// thanhdat.showInfoGiangvien()

	outputData, err := json.Marshal(thanhdat)
	if err != nil {
		fmt.Printf("Error when marshal data: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(outputData))
}
