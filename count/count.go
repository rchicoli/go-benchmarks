package count

import (
	"regexp"
	"strings"
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
