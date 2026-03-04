package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Square struct {
	Width  float64 `json:"width" description:"The width of the square"`
	Height float64 `json:"height" description:"The height of the square"`
}

// Dientich hinh vuong
// - Dientich = chieu dai * chieu rong
// @return float64
func (s *Square) Area() float64 {
	return s.Width * s.Height
}

// Chu vi hinh vuong
// - Chu vi = 2 * (chieu dai + chieu rong)
// @return float64
func (s *Square) Perimeter() float64 {
	return 2 * (s.Width + s.Height)
}

// Lay du lieu tu ban phim
//   - Su dung Scanf
func inputSquare() Square {
	var square Square
	for {
		fmt.Println("Vui long nhap chieu dai hinh chu nhat:")
		_, err := fmt.Scanf("%f", &square.Width)
		if err == nil && square.Width > 0 {
			break
		}
		fmt.Println("Kich thuoc chieu dai phai lon hon 0")
	}

	for {
		fmt.Println("Vui long nhap chieu rong hinh chu nhat:")
		_, err := fmt.Scanf("%f", &square.Height)
		if err == nil && square.Height > 0 {
			break
		}
		fmt.Println("Kich thuoc chieu rong phai lon hon 0")
	}
	return square
}

// Lay du lieu tu ban phim
//   - Su dung bufio va strconv de doc du lieu tu ban phim va chuyen doi sang kieu float64
func inputSquareV2() Square {
	var width, height float64
	for {
		var err error
		width, err = readFloat("Vui long nhap kich thuoc chieu rong hinh chu nhat: ")

		if err != nil || width <= 0 {
			fmt.Println("Kich thuoc chieu rong phai lon hon 0")
		} else {
			break
		}
	}

	for {
		var err error
		height, err = readFloat("Vui long nhap kich thuoc chieu dai hinh chu nhat: ")
		if err != nil || height <= 0 {
			fmt.Println("Kich thuoc chieu dai phai lon hon 0")
		} else {
			break
		}
	}

	return Square{
		Width:  width,
		Height: height,
	}
}

func readFloat(prompt string) (float64, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)
	}

	input = strings.TrimSpace(input)

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting input to float: %v", err)
	}
	return num, nil
}

func main() {
	// square := Square{
	// 	Width:  5,
	// 	Height: 5,
	// }

	// jsonData, err := json.Marshal(square)
	// if err != nil {
	// 	println("Error marshaling square data:", err.Error())
	// 	os.Exit(1)
	// }
	// println("JSON data of square:", string(jsonData))

	// area := square.Area()
	// parameter := square.Perimeter()
	// fmt.Printf("Perimeter of square: %.2f \n", parameter)
	// fmt.Printf("Area of square: %.2f \n", area)
	square := inputSquareV2()
	area := square.Area()
	parameter := square.Perimeter()
	fmt.Printf("Chu vi hinh chu nhat: %.2f \n", parameter)
	fmt.Printf("Dien tich hinh chu nhat: %.2f \n", area)
}
