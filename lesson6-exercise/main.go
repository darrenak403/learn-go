package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fibonacci(n int) {
	term1, term2 := 0, 1

	if n == 1 {
		fmt.Printf("Day so fibonacci: %d \n", term1)
		fmt.Printf("Tong day so fibonancci: %d \n", term1)
	} else if n == 2 {
		fmt.Printf("Day so fibonacci: %d %d \n", term1, term2)
		fmt.Printf("Tong day so fibonancci: %d \n", term1+term2)
	} else {
		total := term1 + term2

		fmt.Printf("Day so fibonacci: %d %d ", term1, term2)

		for i := 3; i <= n; i++ {
			term3 := term1 + term2

			fmt.Printf("%d ", term3)

			total += term3

			term1, term2 = term2, term3
		}

		fmt.Println()
		fmt.Printf("\nTong day so fibonancci: %d \n", total)
	}
}

func tongCacSoTu1DenN(n int) int {
	if n == 1 {
		return 1
	}
	return n + tongCacSoTu1DenN(n-1)
}

func readInt(prompt string) (int, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)
	}

	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("error converting input to integer: %v", err)
	}
	return num, nil
}

func main() {
	for {
		fmt.Println("Chon bai tap:")
		fmt.Println("1. Tong cac so tu 1 den n")
		fmt.Println("2. Fibonacci")
		fmt.Println("0. Thoat")

		var choice int
		for {
			var err error
			choice, err = readInt("Nhap lua chon: ")
			if err != nil || choice <= 0 {
				fmt.Println("Vui long nhap mot so nguyen hop le")
			} else {
				break
			}
		}
		switch choice {
		case 0:
			fmt.Println("Cam on da su dung chuong trinh, Bye!!")
			return
		case 1:
			var num int
			for {
				var err error
				num, err = readInt("Nhap so n: ")
				if err != nil || num <= 0 {
					fmt.Println("Vui long nhap mot so nguyen duong")
				} else {
					break
				}
			}
			result := tongCacSoTu1DenN(num)
			fmt.Printf("Tong cac so tu 1 den %d la: %d\n", num, result)
		case 2:
			var num int
			for {
				var err error
				num, err = readInt("Nhap so n: ")
				if err != nil || num <= 0 {
					fmt.Println("Vui long nhap mot so nguyen duong")
				} else {
					break
				}
			}

			fibonacci(num)
		default:
			fmt.Println("Lua chon khong hop le, vui long chon lai")
		}
	}

}
