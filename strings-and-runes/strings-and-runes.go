package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	const e = "hello"
	fmt.Println("Len s -> สวัสดี:", len(s))
	fmt.Println("Len e -> hello:", len(e))

	fmt.Println("\nIndexing s for hex values of bytes in code points -> สวัสดี:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("\nIndexing e for hex values of bytes in code points -> hello:")
	for i := 0; i < len(e); i++ {
		fmt.Printf("%x ", e[i])
	}
	fmt.Println()

	fmt.Println("\nRune count s -> สวัสดี:", utf8.RuneCountInString(s))
	fmt.Println("Rune count e -> hello:", utf8.RuneCountInString(e))

	fmt.Println("\nDecode rune with offset s -> สวัสดี:")
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nDecode rune with offset e -> hello:")
	for idx, runeValue := range e {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nDecode with DecodeRuneInString s -> สวัสดี:")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

	fmt.Println("\nDecode with DecodeRuneInString e -> hello:")
	for i, w := 0, 0; i < len(e); i += w {
		runeValue, width := utf8.DecodeRuneInString(e[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 'l' {
		fmt.Println("Found el")
	} else if r == 'ส' {
		fmt.Println("Found so sua")
	}
}
