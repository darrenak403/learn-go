package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		num, err := strconv.Atoi(input)
		if err == nil && num > 0 {
			return num
		}
		fmt.Println("❌ Vui long nhap so nguyen duong!")
	}
}

func GetPositiveFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		num, err := strconv.ParseFloat(input, 64)
		if err == nil && num >= 0 {
			return num
		}
		fmt.Println("❌ Vui long nhap so thuc duong!")
	}
}

func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}
		fmt.Println("❌ Vui long nhap du lieu!")
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

type HasId interface {
	GetId() int
}

func IsIdUnique[T HasId](id int, list []T) bool {
	for _, item := range list {
		if item.GetId() == id {
			return false
		}
	}
	return true
}
