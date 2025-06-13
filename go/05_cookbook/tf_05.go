package main

import "unicode"

var data []rune

func FilterCharsAndNormalize() {
	// for i, c := range data {
	// 	switch {
	// 	case c >= 'A' && c <= 'Z':
	// 		data[i] = c + 32
	// 	case c < 'a' || c > 'z':
	// 		data[i] = ' '
	// 	}
	// }
	var filtered []rune
	for _, r := range data {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			filtered = append(filtered, unicode.ToLower(r))
		} else {
			filtered = append(filtered, ' ')
		}
	}
	data = filtered
}
