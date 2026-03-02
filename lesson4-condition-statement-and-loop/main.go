package main

import "fmt"

func main() {
	// Bai 1: Cau lenh dieu kien if-else
	// age := 22
	// if age >= 18 {
	// 	fmt.Println("Ban da du tuoi de bo phieu")
	// } else {
	// 	fmt.Println("Ban chua du tuoi de bo phieu")
	// }

	// Bai 2: Cau lenh dieu kien switch-case
	// day := 3
	// switch day {
	// case 1:
	// 	fmt.Println("Hom nay la thu 2")
	// case 2:
	// 	fmt.Println("Hom nay la thu 3")
	// case 3:
	// 	fmt.Println("Hom nay la thu 4")
	// case 4:
	// 	fmt.Println("Hom nay la thu 5")
	// case 5:
	// 	fmt.Println("Hom nay la thu 6")
	// case 6:
	// 	fmt.Println("Hom nay la thu 7")
	// case 7:
	// 	fmt.Println("Hom nay la chu nhat")
	// default:
	// 	fmt.Println("Ngay khong hop le")
	// }

	// Bai 3: Vong lap for
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Gia tri cua i: %d \n", i)
	// }

	// j := 1
	// for j <= 10 {
	// 	fmt.Printf("Gia tri cua j: %d \n", j)
	// 	j++
	// }

	// for k := 1; k <= 100; k++ {
	// 	if k%9 == 0 {
	// 		fmt.Printf("Gia tri cua k: %d \n", k)
	// 	}
	// }

	// break
	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Printf("Gia tri cua i: %d \n", i)
	// }

	// // continue
	// for j := 1; j <= 10; j++ {
	// 	if j%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Printf("Gia tri cua j: %d \n", j)
	// }

	//Bai tap"
	//Bai1: chay tu 1 - 100 bo qua cac so ngau nghien 6, 48, 75, 89
	// for k := 1; k <= 100; k++ {
	// 	if k == 6 || k == 48 || k == 75 || k == 89 {
	// 		continue
	// 	}
	// 	fmt.Printf("Gia tri cua k: %d \n", k)
	// }

	//Bai2: chay tu 1 - 100 va in ra cac so le, cu 3 so thi xuong dong va ko co dau o cuoi dong, moi so cach nhau bang dau phay, bo dau phay o vi tri cuoi
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 1 {
	// 		fmt.Print(i)
	// 		if (i+1)%6 == 0 {
	// 			fmt.Println()
	// 		} else if i != 99 {
	// 			fmt.Print(", ")
	// 		}
	// 	}
	// }

	//Bai3: xay cdung ung dung hien thi bang cuu chuong, cho phep nguoi nguoi dung nhap so bat dau va ket thuc, kiem tra dieu kien input neu nguoi dung nhap sai, moi bang cuu chuong co tieu de hien thi

	var start, end int
	fmt.Print("Nhap so bat dau: ")
	fmt.Scanln(&start)
	fmt.Print("Nhap so ket thuc: ")
	fmt.Scanln(&end)
	if start < 1 || end > 10 || start > end {
		fmt.Println("Vui long nhap so tu 1 den 10 va so bat dau phai nho hon hoac bang so ket thuc")
		return
	}
	for i := start; i <= end; i++ {
		fmt.Printf("Bang cuu chuong %d: \n", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}
