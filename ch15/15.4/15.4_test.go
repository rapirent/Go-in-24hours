package example03

import "testing"

func TestFrTranslation(t *testing.T) {
    got := translate("fr")
    want := "Bonjour"
    if got != want {
        t.Fatalf("Expected %q, got %q", want, got)
    }
}

func TestUSTranslation(t *testing.T) {
    got := Greeting("George", "en-US")
    want := "Hello George"
    if got != want {
        t.Fatalf("Expected %q, got %q", want, got)
    }
}


type GreetingTest struct {
    name string
    locale string
    want string
}

var greetingTests = []GreetingTest{
    {"George", "en-US", "Hello George"},
    {"Chloe", "fr-FR", "Bonjour Choloe"},
    {"Giuseppe", "it-IT", "Ciao Guiseppe"},
}

func TestGreeting(t *testing.T) {
    for _, test := range greetingTests {
        got := Greeting(test.name, test.locale)
        if got != test.want {
            t.Errorf("Greeting(%s,%s) = %v; want %v", test.name, test.locale, got, test.want)
        }
    }
}
