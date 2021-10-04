package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    flag.Usage = func() {
        usageText := `Usage example04 [OPTION]
An example of customizing usage output

-s, --s xxxxx
-i, --i xxxxx`
        fmt.Fprintf(os.Stderr, "%s\n", usageText)
    }
    s := flag.String("s", "Hello world", "String help text")
    i := flag.Int("i", 1, "Int help text")
    b := flag.Bool("b", false, "Bool help text")
    flag.Parse()
    fmt.Println("value of s:", *s)
    fmt.Println("value of i:", *i)
    fmt.Println("value of b:", *b)
}
