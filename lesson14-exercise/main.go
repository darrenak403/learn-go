package main

import (
	"fmt"

	"darrenak.com/learn-go-basics/student"
	"darrenak.com/learn-go-basics/teacher"
	"darrenak.com/learn-go-basics/utils"
)

func main() {
	for {
		utils.ClearScreen()

		fmt.Println("🎓 CHUONG TRINH QUAN LY 🎓")
		fmt.Println("1. Quan ly sinh vien")
		fmt.Println("2. Quan ly giang vien")
		fmt.Println("3. Thoat chuong trinh")

		choice := utils.GetPositiveInt("👉 Chon chuc nang: ")

		switch choice {
		case 1:
			student.StudentMenu()
		case 2:
			teacher.TeacherMenu()
		case 3:
			return
		default:
			fmt.Println("❌ Lua chon khong hop le!")
		}

		utils.ReadInput("\nNhan phim Enter de tiep tuc...")
	}
}
