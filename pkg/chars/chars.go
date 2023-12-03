package chars

func IsNum(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsLower(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func IsCapital(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func IsLetter(c rune) bool {
	return IsLower(c) || IsCapital(c)
}

func NumVal(c rune) int {
	return int(c - '0')
}
