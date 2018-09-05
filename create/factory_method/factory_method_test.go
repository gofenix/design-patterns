package factory_method

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	fishFactory := FishFactory{}
	fish := fishFactory.CreateAnimal()
	fish.Move(100)

	dogFactory := DogFactory{}
	dog := dogFactory.CreateAnimal()
	dog.Move(200)
}
