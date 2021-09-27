package main

import (
    "fmt"
)

func main() {
    // make a slice
    var cheeses = make([]string, 3)
    cheeses[0] = "Mariolles"
    cheeses[1] = "Epoisses de Bourgogne"
    cheeses[2] = "Camembert"
    fmt.Println(len(cheeses))
    fmt.Println(cheeses)
    cheeses = append(cheeses[:2], cheeses[2+1:]...)
    fmt.Println(len(cheeses))
    fmt.Println(cheeses)
}
