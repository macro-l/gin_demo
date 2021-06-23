package main

import "fmt"

type A interface {
	Foo()
}

type B interface {
	A
	Bar()
}

type T struct{}

func (t T) Foo() {
	fmt.Println("call Foo function from interface A.")
}

func (t T) Bar() {
	fmt.Println("call Bar function from interface B.")
}

type U struct{}

func (u U) Foo() {
	fmt.Println("call Foo function from interface A.")
}

func main() {
	var t = T{}
	var a A = t
	a.Foo()

	// 错误信息： cannot use u (type U) as type B in assignment:
	//var u = U{}
	//var b B = u // error
	//b.Foo()
}
