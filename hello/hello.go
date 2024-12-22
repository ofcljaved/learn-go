package main

import (
    "fmt"
    "log"

    "greetings"
)

func main() {
    log.SetPrefix("greetins: ")
    log.SetFlags(0)

    message, err:= greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}