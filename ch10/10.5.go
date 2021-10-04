package main

import (
    "fmt"
)

func main() {
    fmt.Println("This is executed")
    panic("Oh no. Ican do no more. Goobye.")
    fmt.Println("this is not executed")
}
