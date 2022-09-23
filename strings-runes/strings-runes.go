package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี" // UTF-8 encoded text "sua dee ka"
	// Strings are eqivalent to []byte, the length of the raw bytes
	// will be stored.
	fmt.Println("Length:", len(s))

	//
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // %x %X for format with byte
	}
	// RuneCountInString decode UTF-8 rune sequencially
	fmt.Println("\nRune count:", utf8.RuneCountInString(s))

	// Range loop handles strings! nad decodes each rune
	fmt.Println("1.Range string")
	for idx, runeVal := range s {
		// %#U : add format of U+0078 'x' to the head
		fmt.Printf("%#U starts at %d\n", runeVal, idx) // U+0E2A 'ส' starts at 0
	}

	fmt.Println("\n2.Using DecodeRuneInString, doing the same as 1")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(runes rune) {
	if runes == 't' {
		fmt.Println("found tee")
	} else if runes == 'ส' {
		fmt.Println("found sound sua")
	}
}
