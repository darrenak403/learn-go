package student

import (
	"fmt"
	"strconv"

	"darrenak.com/learn-go-basics/utils"
)

var studentList []Student

func AddStudent() {
	var scores []float64
	fmt.Println("===-=-=== Them sinh vien -=-=-=====")
	id := utils.GetPositiveInt("- Nhap id: ")
	name := utils.GetNonEmptyString("- Nhap ten: ")
	class := utils.GetNonEmptyString("- Nhap lop: ")
	totalScore := utils.GetPositiveInt("- Nhap so luong diem: ")
	for i := 1; i <= totalScore; i++ {
		score := utils.GetPositiveFloat("- Nhap diem thu " + strconv.Itoa(i) + ": ")
		scores = append(scores, score)
	}

	student := Student{
		ID:    id,
		Name:  name,
		Class: class,
		Score: scores,
	}

	studentList = append(studentList, student)

	fmt.Println("✅ Them sinh vien thanh cong")

}

func RemoveStudent() {

}

func UpdateStudent() {

}

func ListStudents() {
	fmt.Println("-=-=-=-=-= Danh sach sinh vien -=-====--=")
	if len(studentList) == 0 {
		fmt.Println("❌ Khong co sinh vien nao")
		return
	}
	for _, student := range studentList {
		fmt.Printf("%s\n", student.GetInfo())
	}
}

func SearchStudent() {

}

func StudentMenu() {
	for {
		utils.ClearScreen()

		fmt.Println("=== QUAN LY SINH VIEN ===")
		fmt.Println("1. Them sinh vien")
		fmt.Println("2. Xoa sinh vien")
		fmt.Println("3. Sua sinh vien")
		fmt.Println("4. Danh sach sinh vien")
		fmt.Println("5. Tim kiem sinh vien")
		fmt.Println("6. Quay lai")

		choice := utils.GetPositiveInt("👉Chon chuc nang: ")

		switch choice {
		case 1:
			AddStudent()
		case 2:
			RemoveStudent()
		case 3:
			UpdateStudent()
		case 4:
			ListStudents()
		case 5:
			SearchStudent()
		case 6:
			return
		default:
			fmt.Println("Lua chon khong hop le!")
		}

		utils.ReadInput("\nNhan phim Enter de tiep tuc...")
	}
}
