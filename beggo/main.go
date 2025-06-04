package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello jaed")
	var name string = "Javed"

	fmt.Printf("This is my name %q\n", name)

	age := 23
	fmt.Printf("Currently aged at %d\n", age)

	var country string
	country = "India"
	fmt.Printf("Living in %s\n", country)

	var (
		jila     = "New Delhi"
		code int = 110015
	)
	fmt.Printf("JIla adhyax of %q, with postal code: %d\n", jila, code)

	const username = "ofcljaved"

	fmt.Printf("The fixed username %q\n", username)

	const (
		Monday  = iota + 1 //1
		Tuesday            //2
		Wednesday
		Thursday
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday)

	var fname, lname = "Javed", "Ansari"
	onePieceName := onepiecify(fname, lname)
	fmt.Println(onePieceName)

	for i := 0; i < len(onePieceName); i++ {
		fmt.Println(onePieceName[i])
	}

	fruits := []string{"apple", "banana", "mango"}
	moreFruits := []string{
		"tomato",
		"cucumber",
	}

	fruits = append(fruits, moreFruits...)
	fmt.Println(fruits)
}

func onepiecify(fname string, lname string) string {
	var op strings.Builder
	op.WriteString(fname)
	op.WriteString(" D ")
	op.WriteString(lname)
	return op.String()
}
