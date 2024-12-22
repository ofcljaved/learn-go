package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) (z float64) {
    z = 1
    threshold := 1e-9
    iterations := 0
    for {
        prevZ := z
        z -=(z*z - x) / (2*z)
        iterations++
        fmt.Printf("z in iteration %v is %v\n", iterations, z)

        if math.Abs(z - prevZ) < threshold {
            break
        }
    }
    return
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println()
    fmt.Println(Sqrt(3))
    fmt.Println()
    fmt.Println(Sqrt(4))
}
