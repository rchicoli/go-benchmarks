package count

import "regexp"

// Match non-space character sequences
var re = regexp.MustCompile(`[\S]+`)

func WordCountRegexp(str string) int {

	// Find all matches and return count
	counter := re.FindAllString(str, -1)

	return len(counter)

}
