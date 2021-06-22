package Interface_Copy

import "fmt"

type Number1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number2 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
	Add(i int)
}

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func (n *Number) Add(i int) {
	*n = *n + Number(i)
}

func main() {
	//var num1 Number = 1
	//var num2 Number1 = num1
	//var num3 Number2 = num2

	var num1 oop1.Number = 1
	var num2 oop2.Number2 = &num1
	var num3 oop1.Number1 = num2

	// 数学集合思想
	// 子集向上兼容
	//var num1 oop1.Number = 1
	//var num2 oop1.Number1 = &num1
	//var num3 oop2.Number2 = num2 // 这一段编译出错

	fmt.Println(num3)

}
