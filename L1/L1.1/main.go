package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name string
	Age  int
}

func NewHuman(name string, age int) *Human {
	return &Human{
		Name: name,
		Age:  age,
	}
}

func (h *Human) Scream(phrase string) {
	scream := strings.ToUpper(phrase)
	fmt.Printf("%s screams: %s\n", h.Name, scream)
}

func (h *Human) Tell(phrase string) {
	fmt.Printf("%s tells: %s\n", h.Name, phrase)
}

func (h *Human) Whisper(phrase string) {
	whisper := strings.ToLower(phrase)
	fmt.Printf("%s whispers: %s\n", h.Name, whisper)
}

type Action struct {
	Human
	ActionType string
}

func NewAction(human *Human, actionType string) *Action {
	return &Action{
		Human:      *human,
		ActionType: actionType,
	}
}

func (a *Action) DoSomething() {
	fmt.Printf("%s does %s\n", a.Name, a.ActionType)
}

func main() {
	human := NewHuman("Bob", 30)
	action := NewAction(human, "work")
	action.DoSomething()
	action.Scream("hi! how are you?")
	action.Tell("Hey! How is it going?")
	action.Whisper("DO YOU HEAR ME?")
	fmt.Printf("%s is %d years old\n", action.Name, action.Age)
}
