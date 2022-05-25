package alphanum

import "testing"

func TestAlphaLower(t *testing.T) {
	alphaLower := AlphaLower()
	if len(alphaLower) != 26 || alphaLower[0] != 'a' || alphaLower[25] != 'z' {
		t.Log("Lowercase alpha characters not correct", alphaLower)
		t.Fail()
	}
}

func TestAlphaUpper(t *testing.T) {
	alphaUpper := AlphaUpper()
	if len(alphaUpper) != 26 || alphaUpper[0] != 'A' || alphaUpper[25] != 'Z' {
		t.Log("Uppercase alpha characters not correct", alphaUpper)
		t.Fail()
	}
}

func TestNumeric(t *testing.T) {
	numeric := Numeric()
	if len(numeric) != 10 || numeric[0] != '0' || numeric[9] != '9' {
		t.Log("Numeric characters not correct", numeric)
		t.Fail()
	}
}

func TestAlphaNumeric(t *testing.T) {
	alphaNum := AlphaNumeric()
	if len(alphaNum) != 62 || alphaNum[0] != 'a' || alphaNum[61] != '9' {
		t.Log("Alphanumeric characters not correct", alphaNum)
		t.Fail()
	}
}
