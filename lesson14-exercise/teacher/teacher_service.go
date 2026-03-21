package teacher

import (
	"fmt"

	"darrenak.com/learn-go-basics/utils"
)

var teacherList []Teacher

func AddTeacher() {
	fmt.Println("===-=-=== Them giang vien -=-=-=====")
	var id int
	for {
		id = utils.GetPositiveInt("- Nhap id: ")
		if utils.IsIdUnique(id, teacherList) {
			break
		}
		fmt.Println("❌ Id da ton tai")
	}
	name := utils.GetNonEmptyString("- Nhap ten: ")
	subject := utils.GetNonEmptyString("- Nhap mon giang day: ")
	baseSalary := utils.GetPositiveFloat("- Nhap luong co ban: ")
	bonus := utils.GetPositiveFloat("- Nhap thuong: ")

	teacher := Teacher{
		Id:         id,
		Name:       name,
		Subject:    subject,
		BaseSalary: baseSalary,
		Bonus:      bonus,
	}

	teacherList = append(teacherList, teacher)

	fmt.Println("✅ Them sinh vien thanh cong")
}

func RemoveTeacher() {

}

func UpdateTeacher() {

}

func ListTeachers() {
	fmt.Println("-=-=-=-=-= Danh sach giang vien -=-====--=")
	for _, teacher := range teacherList {
		fmt.Println(teacher.GetInfo())
	}
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
