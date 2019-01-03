//Package raindrops provides functions to translate numbers into raindrop-speak
package raindrops

import "strconv"

//Convert calculates the raindrop-speak for a given number
func Convert(input int) string {
	output := ""

	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(input)
	}

	return output
}
