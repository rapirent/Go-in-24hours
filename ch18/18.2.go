package main

import (
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hello World\n"))
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":2050", nil)
}
