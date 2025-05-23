package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)

	r := make([]string, 3, 10)
	fmt.Println(r, len(r), cap(r))
	r[0] = "a"
	r[1] = "a"
	r[2] = "a"
	r = append(r, "2")
	fmt.Println(r, len(r), cap(r))
	c := r
	c[2] = "b"
	fmt.Println(c, len(c), cap(c))
	fmt.Println(r, len(r), cap(r))

	d := make([]string, len(r))
	copy(d, r)
	d[2] = "1"
	fmt.Println(d, len(d), cap(d))
	fmt.Println(r, len(r), cap(r))

	l := r[2:4]
	fmt.Println("sl1:", l)
	l[1] = "1"
	fmt.Println(l, len(l), cap(l))
	fmt.Println(r, len(r), cap(r))
}
