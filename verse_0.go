package main

import "fmt"

type Verse_0 struct {
	DefaultVerse
}

func (v Verse_0) Action() string {
	return "Go to the store, buy some more."
}

func (v Verse_0) Beverage() string {
	return "no more bottles"
}

// ! copied exactly from DefaultVerse to get correct dispatch for Verse_0.{Action, Successor}
func (v Verse_0) String() string {
	call := fmt.Sprintf("%s of beer on the wall, %s of beer.", capitalize(v.Beverage()), v.Beverage())
	response := fmt.Sprintf("%s %s of beer on the wall.", v.Action(), capitalize(v.Successor().Beverage()))
	return fmt.Sprintf("%s\n%s", call, response)
}

func (v Verse_0) Successor() Verse {
	return DefaultVerse{99}
}
