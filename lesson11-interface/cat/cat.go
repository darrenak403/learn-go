package cat

import (
	"errors"
	"fmt"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
}

// Create a constructor for Cat
func New(name string) (*Cat, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("cat name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("cat name cannot exceed 50 characters")
	}

	return &Cat{
		Name: name,
	}, nil
}

func (c *Cat) GetName() string {
	return c.Name
}

func (c *Cat) Speak() string {
	return "Meow!"
}

func (c *Cat) Eat() string {
	return fmt.Sprintf("%s is eating fish.", c.Name)
}
