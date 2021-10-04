package main

import (
    "log"
    "os"
)

func main() {
    to, err := os.OpenFile("./deleteme.txt", os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer to.Close()
    err = os.Remove("./deleteme.txt")
    if err != nil {
        log.Fatal(err)
    }
}
