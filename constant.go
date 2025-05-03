package main

import (
	"fmt"
	"math"
)

const s string = "AAND"

func main() {
	fmt.Println(s)

	const n float64 = 55555
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	var b = 3e20 / n
	fmt.Println(b)
	fmt.Println(int64(b))

	fmt.Println(math.Sin(n))
}
