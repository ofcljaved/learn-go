package main

import "fmt"

func main() {
	var a [5]int
	a[4] = 1
	fmt.Println(a)
	fmt.Println(len(a))

	b := [...]string{"ha", 3: "ba", "fao"}
	fmt.Println(b)

	var _2D [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			_2D[i][j] = i + j
		}
	}
	fmt.Println(_2D)
}
