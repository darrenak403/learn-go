package student

import "fmt"

type Student struct {
	Id    int
	Name  string
	Class string
	Score []float64
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("ID: %d | Name: %s | Class: %s | Average Score: %.2f", s.Id, s.Name, s.Class, s.CaculateAverageScore())
}

func (s Student) CaculateAverageScore() float64 {
	if len(s.Score) == 0 {
		return 0
	}
	total := 0.0
	for _, score := range s.Score {
		total += score
	}
	return total / float64(len(s.Score))
}

func (s Student) GetId() int {
	return s.Id
}
