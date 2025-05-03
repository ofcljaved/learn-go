package main

import "fmt"

var a = "declare variable"

func main() {
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	var f bool
	var g string

	//Zero valued variables -> Those without initialization
	fmt.Println(e, f, g)

	h := "hero" // Only inside the function
	fmt.Println(h)
}
