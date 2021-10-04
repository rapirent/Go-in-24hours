package main

import (
    "fmt"
    "time"
)

func slowFunc() {
    time.Sleep(time.Second * 2)
    fmt.Println("sleeper() finished")
}

func main() {
    // non blocking
    go slowFunc()
    fmt.Println("I am now shown straightaway!")
    time.Sleep(time.Second * 3)
}
