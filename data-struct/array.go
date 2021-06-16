package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [3]int
	fmt.Println(reflect.TypeOf(a))
}
