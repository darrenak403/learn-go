package mouse

import (
	"errors"
	"strings"

	"darrenak.com/learn-go-basics/service"
)

type Mouse struct {
	Name string `json:"name"`
}

// Create a constructor for Mouse
func New(name string) (service.AnimalAction, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("Mouse name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("Mouse name cannot exceed 50 characters")
	}

	return &Mouse{
		Name: name,
	}, nil
}

func (m *Mouse) GetName() string {
	return m.Name
}

func (m *Mouse) Speak() string {
	return "Chit chits!"
}

func (m *Mouse) Run() string {
	return "The mouse is running!"
}
