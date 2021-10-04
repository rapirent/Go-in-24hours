// Package example03 shows how to use the godoc tool
package example03

import (
    "errors"
)

type Animal struct {
    Name string
    Age int
}

var ErrNotAnAnimal = errors.New("Name is not an animal")

//Hello sends a greeting to the animal
func (a Animal) Hello() (string, error) {
    if a.Name == "Human" {
        return "", ErrNotAnAnimal
    }
    s := "Hello " + a.Name
    return s, nil
}
