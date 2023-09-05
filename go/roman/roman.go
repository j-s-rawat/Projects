package main

var romanDigitToInt = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func RomanToIntString(input string) (out int) {
	l := len(input)
	for i := 0; i < l; i++ {
		if i+1 < l {
			if romanDigitToInt[string(input[i])] >= romanDigitToInt[string(input[i+1])] {
				out += romanDigitToInt[string(input[i])]
			} else {
				out += romanDigitToInt[string(input[i+1])] - 1
				i++
			}

		} else {
			out += romanDigitToInt[string(input[i])]
		}

	}

	return
}

var romanIntToDigit = []struct {
	I     int
	Roman string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IV"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func IntToRoman(input int) (out string) {
	for _, v := range romanIntToDigit {
		for v.I <= input {
			out += v.Roman
			input -= v.I
		}

	}
	return
}
