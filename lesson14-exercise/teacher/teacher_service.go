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
	fmt.Println("===-=-=== Xoa giang vien -=-=-=====")
	id := utils.GetPositiveInt("- Nhap id: ")
	// 0 1 2 3 4
	// 1 2 3 4 5
	// 1 2 3 4 5
	// i = 2
	// teacherList[:2] = [1, 2]
	// teacherList[3:] = [4, 5]

	for i, teacher := range teacherList {
		if teacher.Id == id {
			teacherList = append(teacherList[:i], teacherList[i+1:]...)
			fmt.Println("✅ Xoa giang vien thanh cong")
			return
		}
	}
	fmt.Println("❌ Khong tim thay giang vien")
}

func UpdateTeacher() {
	fmt.Println("===-=-=== Sua giang vien -=-=-=====")
	id := utils.GetPositiveInt("- Nhap id: ")

	for i, teacher := range teacherList {
		if teacher.Id == id {
			fmt.Println("Nhap thong tin moi (Nhan enter de giu nguyen gia tri hien tai):")
			name := utils.GetOptionalString(fmt.Sprintf("Nhap ten (%s): ", teacher.Name), teacher.Name)
			subject := utils.GetOptionalString(fmt.Sprintf("Nhap mon giang day (%s): ", teacher.Subject), teacher.Subject)
			baseSalary := utils.GetOptionalPositiveFloat(fmt.Sprintf("Nhap luong co ban (%f): ", teacher.BaseSalary), teacher.BaseSalary)
			bonus := utils.GetOptionalPositiveFloat(fmt.Sprintf("Nhap thuong (%f): ", teacher.Bonus), teacher.Bonus)

			teacherList[i] = Teacher{
				Id:         id,
				Name:       name,
				Subject:    subject,
				BaseSalary: baseSalary,
				Bonus:      bonus,
			}

			fmt.Println("✅ Sua giang vien thanh cong")
			return
		}
	}
	fmt.Println("❌ Khong tim thay giang vien")

}

func ListTeachers() {
	fmt.Println("-=-=-=-=-= Danh sach giang vien -=-====--=")
	for _, teacher := range teacherList {
		fmt.Println(teacher.GetInfo())
	}
}

func SearchTeacher() {
	fmt.Println("===-=-=== Tim kiem giang vien -=-=-=====")
	id := utils.GetPositiveInt("- Nhap id: ")

	for _, t := range teacherList {
		if t.Id == id {
			fmt.Println("Tim thay giang vien:")
			fmt.Println(t.GetInfo())
			return
		}
	}
	fmt.Println("❌ Khong tim thay giang vien")
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
