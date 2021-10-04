package main

import (
    "fmt"
    "log"
    "regexp"
)

func main() {
    needle := "chocolate"
    haystack := "Chocolate is my favorite!"
    match, err := regexp.MatchString(needle, haystack)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(match)
    needle = "(?i)chocolate"
    match, err = regexp.MatchString(needle, haystack)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(match)
}
