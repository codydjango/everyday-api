/*
Helper functions.
*/

package main

// GetOrdinal returns ordinals for integers
func GetOrdinal(i int) string {
	j := i % 10
	k := i % 100

	if j == 1 && k != 11 {
		return "st"
	}

	if j == 2 && k != 12 {
		return "nd"
	}

	if j == 3 && k != 13 {
		return "rd"
	}

	return "th"
}
