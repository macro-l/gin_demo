package main

import "fmt"

func main() {
	a := 1 << 20
	a = a / 1024
	fmt.Println(a)
}
