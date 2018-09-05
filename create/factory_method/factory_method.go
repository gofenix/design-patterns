package factory_method

import "fmt"

type Action interface {
	Move(n int)
}

// 父类
type animal struct {
	Name string
}

// fish 使用组合来模拟继承
type fish struct {
	animal
}

// fish 实现接口 Action
func (f fish) Move(n int) {
	fmt.Println(f.Name, n)
}

// dog 使用组合来模拟继承
type dog struct {
	animal
}

// dog 实现接口 Action
func (d dog) Move(n int) {
	fmt.Println(d.Name, n)
}

// 动物工厂
type AnimalFactory interface {
	CreateAnimal() Action
}

// fish工厂
type FishFactory struct {
}

func (fishFactory FishFactory) CreateAnimal() Action {
	return fish{animal{Name: "fish"}}
}

// dog工厂
type DogFactory struct {
}

func (dogFactory DogFactory) CreateAnimal() Action {
	return dog{animal{Name: "dog"}}
}
