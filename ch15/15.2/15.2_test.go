package example02

import "testing"

func TestGreeting(t *testing.T) {
    got := Greeting("George")
    want := "Hello george"
    if got != want {
        t.Fatalf("Expected %q, got %q", want, got)
    }
}
