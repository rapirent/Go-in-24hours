package main

import "fmt"
import "strconv"

func main() {
    var i int = 1
    var a string = " egg"
    intToString := strconv.Itoa(i)
    var breakfast string = intToString + a
    fmt.Println(breakfast)
}
