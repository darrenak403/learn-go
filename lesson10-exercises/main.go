package main

import "fmt"

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
	square := inputSquare()
	area := square.Area()
	parameter := square.Perimeter()
	fmt.Printf("Chu vi hinh chu nhat: %.2f \n", parameter)
	fmt.Printf("Dien tich hinh chu nhat: %.2f \n", area)
}
