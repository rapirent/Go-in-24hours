package main

import (
    "fmt"
)

func main() {
    // make a slice
    var cheeses = make([]string, 2)
    cheeses[0] = "Mariolles"
    cheeses[1] = "Epoisses de Bourgogne"
    cheeses = append(cheeses, "Camembert", "Reblochon", "Picodon")
    fmt.Println(cheeses)
}
