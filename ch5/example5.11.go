package main

import (
    "fmt"
)

func main() {
    // the defer statement will be executed in the reverse order of their presence
    defer fmt.Println("I am the first defer statement")
    defer fmt.Println("I am the second defer statement")
    defer fmt.Println("I am the third defer statement")
    fmt.Println("Hello World")
}
