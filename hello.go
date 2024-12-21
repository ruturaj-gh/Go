package main

import (
    "Go/hello/greetings"
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println("hello world")
    fmt.Println(quote.Go())
    message := greetings.Hello("Ruturaj")
    fmt.Println(message)
}
