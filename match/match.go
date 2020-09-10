package match

func IsMatch(text string, pattern string) bool {
	if len(pattern) == 0 {
		return len(text) <= 0
	}
	first := len(text) > 0 && (pattern[0] == text[0] || pattern[0] == '.')
	if len(pattern) >= 2 && pattern[1] == '*' {
		return IsMatch(text, pattern[2:]) || (first && IsMatch(text[1:], pattern))
	} else {
		return first && IsMatch(text[1:], pattern[1:])
	}
}