package main
import (
    "fmt"
    "log"
    "time"
)

func main() {
    s1 := "2007-09-03T18:00:00+00:00"
    s2 := "2007-09-04T18:00:00+00:00"
    today, err := time.Parse(time.RFC3339, s1)
    if err != nil {
        log.Fatal(err)
    }
    tomorrow, err := time.Parse(time.RFC3339, s2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(today.After(tomorrow))
    fmt.Println(today.Before(tomorrow))
    fmt.Println(today.Equal(tomorrow))
    nt := today.Add(2 * time.Minute)
    subTime := time.Date(2007, 9, 02, 18, 0, 0, 0, time.UTC)
    nt2 := today.Sub(subTime)
    nt3 := today.Add(-24 * time.Hour)
    fmt.Println(nt)
    fmt.Println(nt2)
    fmt.Println(nt3)
}
