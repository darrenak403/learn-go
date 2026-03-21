package teacher

import "fmt"

type Teacher struct {
	Id         int
	Name       string
	Subject    string
	BaseSalary float64
	Bonus      float64
}

func (t Teacher) GetInfo() string {
	return fmt.Sprintf("ID: %d | Name: %s | Subject: %s | Base Salary: %.2f | Bonus: %.2f | Total Salary: %.2f", t.Id, t.Name, t.Subject, t.BaseSalary, t.Bonus, t.CalculateTotalSalary())
}

func (t Teacher) CalculateTotalSalary() float64 {
	return t.BaseSalary + t.Bonus
}

func (s Teacher) GetId() int {
	return s.Id
}
