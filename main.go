package main

import (
	"fmt"
)

func main() {
	fmt.Println("ok")
}

// Verse returns verse for `bottles` bottles in the 99 Bottles of Beer on the Wall
func Verse(bottles int) string {
	switch bottles {
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
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\n"+
			"Take one down, pass it around. %d bottles of beer on the wall.",
			bottles, bottles, bottles-1)
	}
}
