// Reverse words and sentences
//
// These simple implementations have many problems with:
//  Strings that are not NUL terminated
//  Variable length encodings such as UTF8 or UTF16 etc
//  Ideographic languages such as Chinese or Korean etc
//  Languages with conjoining characters such as Arabic etc
//  Word separators other than space such as tab, nbsp, punctuation characters

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Reverse a string, naively
// This is not safe for multi-byte characters i.e. other than simple ASCII
func reverse_naive(str string) (rev string) {
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}

// Reverse a string, unicode aware
// This is not safe for conjoining characters e.g. ae or languages such as Arabic
func reverse(str string) (rev string) {
	runes := make([]rune, utf8.RuneCountInString(str))

	var size int
	for i := 0; i < len(runes) && len(str) > 0; i++ {
		runes[i], size = utf8.DecodeLastRuneInString(str)
		str = str[:len(str)-size]
	}
	return string(runes)
}

// Reverse a sentence
// This does not properly handle ,<sp> or .<sp> etc
func reverse_sentence(str string) (rev string) {
	// Find each word in the sentence and reverse it
	start := -1
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		if unicode.IsMark(r) || unicode.IsPunct(r) || unicode.IsSpace(r) || unicode.IsSymbol(r) {
			if start != -1 {
				// Found the end of a word, reverse it
				rev += reverse(str[start:i])

				// Look for a new word
				start = -1
			}
			// Keep the symbol
			rev += string(r)
		} else {
			if start == -1 {
				// Found the start of a word
				start = i
			}
		}

		// Move on to the next rune
		i += size
	}

	// Reverse all of the characters in the sentence
	rev = reverse(rev)
	return rev
}

func main() {
	str := "Hello, 世界"
	sentence := "Hello, 世界 - reverse me."

	fmt.Printf("Reverse a string, naively:\n%s -> %s\n\n", str, reverse_naive(str))
	fmt.Printf("Reverse a string, unicode aware:\n%s -> %s\n\n", str, reverse(str))
	fmt.Printf("Reverse a sentence:\n%s -> %s\n\n", sentence, reverse_sentence(sentence))
}
