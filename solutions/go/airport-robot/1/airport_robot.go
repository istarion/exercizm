package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(string) string
}

func SayHello(name string, g Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

type Italian struct {}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {}

func (i Portuguese) LanguageName() string {
    return "Portuguese"
}

func (i Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}