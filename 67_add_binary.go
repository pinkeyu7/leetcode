package main

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1

	res := ""
	carry := 0

	for i >= 0 || j >= 0 {
		temp := carry

		if i >= 0 && string(a[i]) == "1" {
			temp++
		}

		if j >= 0 && string(b[j]) == "1" {
			temp++
		}

		carry = temp / 2
		temp = temp % 2

		if temp%2 == 1 {
			res = "1" + res
		} else {
			res = "0" + res
		}

		i--
		j--
	}

	if carry == 1 {
		res = "1" + res
	}

	return res
}
