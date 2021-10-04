/**
javascript mapping
Boolean <-> bool
Number <-> float64
String <-> string
Array <-> []interface{}
Object <-> map[string]interface{}
Null <-> nil
**/
package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Switch struct {
    On bool `json:"on"`
}

func main() {
    jsonStringData := `{"on":"true"}`
    jsonByteData := []byte(jsonStringData)
    p := Switch{}
    err := json.Unmarshal(jsonByteData, &p)
    // it will be rised, because "true" is a string, can't be converted to boolean
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", p)
}
