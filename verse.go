package main

import (
	"fmt"
	"unicode"
)

type Verse interface {
	Action() string
	Beverage() string
	String() string
	Successor() Verse
}

func GetVerse(bottles int) Verse {
	switch bottles {
	case 1:
		return Verse_1{DefaultVerse{bottles}}
	case 0:
		return Verse_0{DefaultVerse{bottles}}
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
	return fmt.Sprintf("%d bottles", v.bottles)
}

func (v DefaultVerse) String() string {
	call := fmt.Sprintf("%s of beer on the wall, %s of beer.", capitalize(v.Beverage()), v.Beverage())
	response := fmt.Sprintf("%s %s of beer on the wall.", v.Action(), capitalize(v.Successor().Beverage()))
	return fmt.Sprintf("%s\n%s", call, response)
}

func (v DefaultVerse) Successor() Verse {
	return GetVerse(v.bottles - 1)
}

func capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}
