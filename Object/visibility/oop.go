package main

import (
	"fmt"
	"visibility/animal"
)

func main() {
	ani := animal.NewAnimal("狗")
	dog := animal.Dog{ani}
	fmt.Println(dog.GetName(), dog.Call(), dog.FavorFood())

	// 不可见的所属类调用
	fmt.Println(dog.Animal.Call())
	// 不可见的外部调用
	// 错误提示： dog.Animal.happy undefined (cannot refer to unexported field or method animal.Animal.happy)
	//fmt.Println(dog.Animal.happy())

	// 额外测试下继承
	fmt.Println(dog.Animal.Unhappy())
}
