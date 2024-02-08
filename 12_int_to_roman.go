package main

func intToRoman(num int) string {
	result := ""

	thousand := num / 1000
	hundred := num % 1000 / 100
	ten := num % 100 / 10
	one := num % 10

	for i := 0; i < thousand; i++ {
		result += "M"
	}

	if hundred == 9 {
		result += "CM"
	} else if hundred >= 5 {
		result += "D"
		for i := 0; i < hundred-5; i++ {
			result += "C"
		}
	} else if hundred == 4 {
		result += "CD"
	} else {
		for i := 0; i < hundred; i++ {
			result += "C"
		}
	}

	if ten == 9 {
		result += "XC"
	} else if ten >= 5 {
		result += "L"
		for i := 0; i < ten-5; i++ {
			result += "X"
		}
	} else if ten == 4 {
		result += "XL"
	} else {
		for i := 0; i < ten; i++ {
			result += "X"
		}
	}

	if one == 9 {
		result += "IX"
	} else if one >= 5 {
		result += "V"
		for i := 0; i < one-5; i++ {
			result += "I"
		}
	} else if one == 4 {
		result += "IV"
	} else {
		for i := 0; i < one; i++ {
			result += "I"
		}
	}

	return result
}
