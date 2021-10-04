package main

import (
    "fmt"
)

func main() {
    a := `After a backslash, certain single character escapes represent special values
n is a line feed or new line
    \t t is a tab`
    fmt.Println(a)
}
