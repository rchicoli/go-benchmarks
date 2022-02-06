package count

import (
	"regexp"
	"strings"
	"unicode"
)

// Match non-space character sequences
var re = regexp.MustCompile(`[\S]+`)

func WordCountRegexp(str string) int {

	// Find all matches and return count
	counter := re.FindAllString(str, -1)

	return len(counter)

}

func WordCount(str string) int {
	words := strings.Fields(str)
	counter := 0
	for range words {
		counter++
	}
	return counter
}

func WordCountUnicode(str string) int {
	counter := 0
	for _, v := range strings.TrimSpace(str) {
		if unicode.IsSpace(v) {
			counter++
		}
	}
	return counter + 1
}
