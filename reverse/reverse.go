package reverse

func Reverse(str string) string {
	var revStr string
	for i := len(str) - 1; i >= 0; i-- {
		revStr = revStr + string(str[i])
	}
	return revStr
}

func ReverseUnicode(s string) string {
	unicode := []rune(s)
	l := len(unicode)
	rev := make([]rune, l)
	l--
	for k, v := range unicode {
		rev[l-k] = v
	}
	return string(rev)
}

func ReverseRange(str string) string {
	var result string
	for _, v := range str {
		result = string(v) + result
	}
	return result
}
