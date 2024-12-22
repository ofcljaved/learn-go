package main

import "fmt"

func main() {
    table(2)
    sum_count()
}

func table(num int) {
    for i:=1; i <= 10; i++ {
        fmt.Printf("%v X %v = %v\n", num, i, num * i)
    }
}

func sum_count() {
    iteration_count := 0
    sum := 1
    for sum < 1000 {
        sum += sum
        iteration_count++ 
    }

    fmt.Println(`The number of time loop run to calcualute `,sum,` is`, iteration_count)
}
