package count

import "regexp"

func WordCountRegexp(str string) int {
	// Match non-space character sequences
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count
	counter := re.FindAllString(str, -1)

	return len(counter)

}
