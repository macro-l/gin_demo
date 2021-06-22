package Interface

import "fmt"

type Integer int

// 归属于 Integer类下， 在指针引用时，会在*Integer类自动生成一个新的与之对于的指针成员方法
func (a Integer) Equal(i Integer) bool {
	return a == i
}

// 归属于 Integer类下， 在指针引用时，会在*Integer类自动生成一个新的与之对于的指针成员方法
func (a Integer) LessThan(i Integer) bool {
	return a < i
}

// 归属于 Integer类下， 在指针引用时，会在*Integer类自动生成一个新的与之对于的指针成员方法
func (a Integer) MoreThan(i Integer) bool {
	return a > i
}

// 归属于 Integer类的*Integer 属性下
func (a *Integer) Increase(i Integer) {
	*a = *a + i
}

// 归属于 Integer类的*Integer 属性下
func (a *Integer) Decrease(i Integer) {
	*a = *a - i
}

type IntNumber interface {
	Equal(i Integer) bool
	LessThan(i Integer) bool
	MoreThan(i Integer) bool
	Increase(i Integer)
	Decrease(i Integer)
}

type IntNumber2 interface {
	Equal(i Integer) bool
	LessThan(i Integer) bool
	MoreThan(i Integer) bool
}

func main() {
	var a Integer = 1
	var b IntNumber = &a

	// Integer里 Increase和Decrease属于*Integer
	// 所以值传递时interface无法访问Increase和Decrease
	// 错误信息： cannot use a (type Integer) as type IntNumber in assignment:
	//var bb IntNumber = a

	be := b.Equal(1)
	fmt.Println(be)

	b.Increase(1)
	fmt.Println(a)

	var c IntNumber2 = a
	ce := c.Equal(1)
	fmt.Println(ce)
}
