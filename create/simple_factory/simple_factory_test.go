package simple_factory

import (
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	fish := NewAnimalFactory().CreateAnimal("fish")
	fish.Move(100)

	dog := NewAnimalFactory().CreateAnimal("dog")
	dog.Move(200)
}
