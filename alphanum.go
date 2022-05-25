package alphanum

func AlphaLower() []rune {
	var i rune
	lower := make([]rune, 26)
	for i = 0; i < 26; i++ {
		lower[i] = 97 + i
	}
	return lower
}

func AlphaUpper() []rune {
	var i rune
	upper := make([]rune, 26)
	for i = 0; i < 26; i++ {
		upper[i] = 65 + i
	}
	return upper
}

func Numeric() []rune {
	var i rune
	numeric := make([]rune, 10)
	for i = 0; i < 10; i++ {
		numeric[i] = 48 + i
	}
	return numeric
}

func AlphaNumeric() []rune {
	return append(append(AlphaLower(), AlphaUpper()...), Numeric()...)
}
