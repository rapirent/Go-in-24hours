package main

import "fmt"
import "errors"

func main() {
    err := errors.New("Something went wrong")
    if err != nil {
        fmt.Println(err)
    }
}
