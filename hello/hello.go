package main

import (
    "fmt"
    "log"

    "greetings"
)

func main() {
    log.SetPrefix("greetins: ")
    log.SetFlags(0)

    names :=[]string{"Jack", "Javed", "Jia"}

    messages, err:= greetings.Hellos(names)

    message, err:= greetings.Hello("Jack")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
    fmt.Println("\r\nMultiple greetings on the way!!!\r\n")
    fmt.Println(messages)
}
