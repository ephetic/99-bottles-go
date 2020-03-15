package main

import (
	"fmt"
)

type Verse interface {
	Action() string
	Beverage() string
	Container() string
	Pronoun() string
	String() string
	Successor() Verse
}

func GetVerse(bottles int) Verse {
	switch bottles {
	case 1:
		return Verse_1{DefaultVerse{bottles: bottles}}
	case 0:
		return Verse_0{DefaultVerse{bottles: bottles}}
	default:
		return DefaultVerse{bottles}
	}
}

type DefaultVerse struct {
	bottles int
}

func (v DefaultVerse) Action() string {
	return "Take one down, pass it around."
}

func (v DefaultVerse) Beverage() string {
	return fmt.Sprintf("%d %s", v.bottles, v.Container())
}

func (v DefaultVerse) Container() string {
	return "bottles"
}

func (v DefaultVerse) Pronoun() string {
	return fmt.Sprint(v.bottles)
}

func (v DefaultVerse) String() string {
	call := fmt.Sprintf("%s of beer on the wall, %s of beer.", capitalize(v.Beverage()), v.Beverage())
	response := fmt.Sprintf("%s %s of beer on the wall.", v.Action(), capitalize(v.Successor().Beverage()))
	return fmt.Sprintf("%s\n%s", call, response)
}

func (v DefaultVerse) Successor() Verse {
	return GetVerse(v.bottles - 1)
}
