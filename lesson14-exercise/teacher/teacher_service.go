package teacher

import (
	"fmt"

	"darrenak.com/learn-go-basics/utils"
)

func AddTeacher() {

}

func RemoveTeacher() {

}

func UpdateTeacher() {

}

func ListTeachers() {

}

func SearchTeacher() {

}

func TeacherMenu() {
	for {
		utils.ClearScreen()

		fmt.Println("=== QUAN LY GIANG VIEN ===")
		fmt.Println("1. Them giang vien")
		fmt.Println("2. Xoa giang vien")
		fmt.Println("3. Sua giang vien")
		fmt.Println("4. Danh sach giang vien")
		fmt.Println("5. Tim kiem giang vien")
		fmt.Println("6. Quay lai")

		choice := utils.GetPositiveInt("👉Chon chuc nang: ")

		switch choice {
		case 1:
			AddTeacher()
		case 2:
			RemoveTeacher()
		case 3:
			UpdateTeacher()
		case 4:
			ListTeachers()
		case 5:
			SearchTeacher()
		case 6:
			return
		default:
			fmt.Println("Lua chon khong hop le!")
		}

		utils.ReadInput("\nNhan phim Enter de tiep tuc...")
	}
}
