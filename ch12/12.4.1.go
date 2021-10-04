package main

import (
    "fmt"
    "time"
)

func pinger(c chan string) {
    t := time.NewTicker(1 * time.Second)
    for {
        c <- "ping"
        <-t.C
    }
}

func main() {
    messages := make(chan string)
    go pinger(messages)
    for i := 0; i< 5; i ++ {
        // <- 在通道左邊代表只讀、<-在通道右邊代表只寫、沒有<-代表可讀寫
        msg := <-messages
        fmt.Println(msg)
    }
}
