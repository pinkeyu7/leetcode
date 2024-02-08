package main

import (
	"strings"
)

func romanToInt(s string) int {
	sum := 0

	charArray := strings.Split(s, "")

	for i := 0; i < len(charArray); i++ {
		if charArray[i] == "M" {
			sum += 1000
		} else if charArray[i] == "D" {
			sum += 500
		} else if charArray[i] == "C" {
			if i < len(charArray)-1 && charArray[i+1] == "D" {
				sum += 400
				i++
			} else if i < len(charArray)-1 && charArray[i+1] == "M" {
				sum += 900
				i++
			} else {
				sum += 100
			}
		} else if charArray[i] == "L" {
			sum += 50
		} else if charArray[i] == "X" {
			if i < len(charArray)-1 && charArray[i+1] == "L" {
				sum += 40
				i++
			} else if i < len(charArray)-1 && charArray[i+1] == "C" {
				sum += 90
				i++
			} else {
				sum += 10
			}
		} else if charArray[i] == "V" {
			sum += 5
		} else if charArray[i] == "I" {
			if i < len(charArray)-1 && charArray[i+1] == "V" {
				sum += 4
				i++
			} else if i < len(charArray)-1 && charArray[i+1] == "X" {
				sum += 9
				i++
			} else {
				sum += 1
			}
		}
	}

	return sum
}
