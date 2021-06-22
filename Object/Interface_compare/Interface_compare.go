package main

import (
	"fmt"
)

/**
动物
*/
type Animal struct {
	name string
}

func NewAnimal(name string) *Animal {
	return &Animal{name: name}
}

func (a Animal) Call() string {
	return "动物的叫声..."
}

func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}

func (a Animal) GetName() string {
	return a.name
}

/**
狗狗
*/
type Dog struct {
	*Animal
}

func (d Dog) FavorFood() string {
	return "骨头"
}

func (d Dog) Call() string {
	return "汪汪汪"
}

type IAnimal interface {
	GetName() string
	Call() string
	FavorFood() string
}

func main() {
	var animal = Animal{"小狗"}
	var ianimal IAnimal = Dog{&animal}
	var ianimal2 IAnimal = animal
	if dog, ok := ianimal.(Dog); ok {
		fmt.Println(dog.GetName())
	} else {
		fmt.Println(ok)
	}
	fmt.Println("NEXT")

	if dog, ok := ianimal2.(Dog); ok {
		fmt.Println(dog.GetName())
	} else {
		fmt.Println(ok)
	}

}
