package main

import "fmt"

type NonDefaultVerse struct {
	DefaultVerse
}

// func (v NonDefaultVerse) Container() string {
// 	return "bottles"
// }

// func (v NonDefaultVerse) Pronoun() string {
// 	return fmt.Sprint(v.bottles)
// }

func (v NonDefaultVerse) String() string {
	switch v.bottles {
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\n" +
			"Take one down, pass it around. 1 bottle of beer on the wall."
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\n" +
			"Take it down, pass it around. No more bottles of beer on the wall."
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\n" +
			"Go to the store, buy some more. 99 bottles of beer on the wall."
	default:
		panic(fmt.Sprintf("Unknown NonDefaultVerse for %d", v.bottles))
	}
}

// func (v NonDefaultVerse) Successor() string {
// 	return fmt.Sprint(v.bottles - 1)
// }
