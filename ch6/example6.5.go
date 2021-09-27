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
    var smellyCheeses = make([]string, 2)
    copy(smellyCheeses, cheeses)
    fmt.Println(smellyCheeses)
}
