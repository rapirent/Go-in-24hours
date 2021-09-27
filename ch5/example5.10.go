package main

import (
    "fmt"
)

func main() {
    // defer means the execution the function when the parent function return
    defer fmt.Println("I am run after the function completes")
    fmt.Println("Hello World")
}
