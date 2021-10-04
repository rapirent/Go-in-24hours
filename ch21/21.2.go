package main

import (
    "io/ioutil"
    "log"
)

func main() {
    b := make([]byte, 0)
    // set the previlege to 0644
    err := ioutil.WriteFile("example02.txt", b, 0644)
    if err != nil {
        log.Fatal(err)
    }
}
