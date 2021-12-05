package main

import (
	"fmt"
	"math"
)

func intLen(x int) int {
	var intLen int
	for x >= 10 {
		x = x / 10
		intLen++
	}
	return intLen
}

func intToDigits(x int, l int) []int {
	var fwList []int
	var d int
	for x != 0 {
		d = x / int(math.Pow(float64(10), float64(l)))
		x = x % int(math.Pow(float64(10), float64(l)))
		l--
		fwList = append(fwList, d)
	}
	fwList = append(fwList, x)
	return fwList
}

func isPalindrome(x int) bool {
	var l int
	var d int
	var digits []int
	if x < 0 {
		return false
	}
	l = intLen(x)
	fmt.Println(l)
	digits = (intToDigits(x, l))
	fmt.Println(digits)
	for x != 0 {
		d = x % 10
		x = x / 10
		if d != digits[0] {
			return false
		}
		digits = digits[1:]
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(1000))
}
