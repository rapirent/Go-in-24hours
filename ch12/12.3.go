package main

import (
    "fmt"
    "time"
)

func receiver(c chan string) {
    for msg := range c {
        fmt.Println(msg)
    }
}

func main() {
    messages := make(chan string, 2)
    messages <- "hello"
    messages <- "world"
    close(messages) // close the channel
    fmt.Println("PUshed two messages onto Channel with no receivers")
    time.Sleep(time.Second * 1)
    receiver(messages)
}
