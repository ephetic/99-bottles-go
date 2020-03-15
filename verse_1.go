package main

import "fmt"

type Verse_1 struct {
	DefaultVerse
}

func (v Verse_1) Action() string {
	return "Take it down, pass it around."
}

// ! copied exactly from DefaultVerse to get correct dispatch for Verse_1.Container
func (v Verse_1) Beverage() string {
	return fmt.Sprintf("%d %s", v.bottles, v.Container())
}

func (v Verse_1) Container() string {
	return "bottle"
}

// ! copied exactly from DefaultVerse to get correct dispatch for Verse_1.{Action, Beverage}
func (v Verse_1) String() string {
	call := fmt.Sprintf("%s of beer on the wall, %s of beer.", capitalize(v.Beverage()), v.Beverage())
	response := fmt.Sprintf("%s %s of beer on the wall.", v.Action(), capitalize(v.Successor().Beverage()))
	return fmt.Sprintf("%s\n%s", call, response)
}
