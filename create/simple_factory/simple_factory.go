package simple_factory

import (
	"fmt"
)

type AnimalFactory struct {
}

func (animalFactory AnimalFactory) CreateAnimal(name string) Action {
	switch name {
	case "fish":
		return fish{animal{Name: "fish"}}
	case "dog":
		return dog{animal{Name: "dog"}}
	default:
		return nil
	}
}

func NewAnimalFactory() AnimalFactory {
	return AnimalFactory{}
}

type Action interface {
	Move(n int)
}

type animal struct {
	Name string
}

type fish struct {
	animal
}

func (f fish) Move(n int) {
	fmt.Println(f.Name, n)
}

type dog struct {
	animal
}

func (d dog) Move(n int) {
	fmt.Println(d.Name, n)
}
