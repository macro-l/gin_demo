package main

import (
	"fmt"
	"reflect"
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

	fmt.Println("ianimal type :", reflect.TypeOf(ianimal))
	fmt.Println("ianimal2 type :", reflect.TypeOf(ianimal2))

	//var a = 1
	//fmt.Println("a type :", a.(type)) // error
	//var a interface{}
	//a = 1
	//fmt.Println("a type :", a.(type)) // error
	myPrintf(1, "2")
}

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case bool:
			fmt.Println(arg, "is a bool value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

//func myPrintf2(args interface{}) {
//	fmt.Println(args.(type))     // error
//}
