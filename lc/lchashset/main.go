package main

import "fmt"

//A happy number is a number defined by the following process:

//Starting with any positive integer, replace the number by the sum of the squares of its digits.
//Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
//Those numbers for which this process ends in 1 are happy.
//Return true if n is a happy number, and false if not.

func sumOfDigits(n int) int {
	var d int
	var totalSum int
	for n != 0 {
		d = n % 10
		n = n / 10
		totalSum = totalSum + (d * d)
	}
	return totalSum
}

func isHappy(n int) bool {
	mySet := make(map[int]bool)
	for {
		n = sumOfDigits(n)
		fmt.Printf("%v\n\n", n)
		if mySet[n] {
			print("loop detected\n")
			return false
		}
		mySet[n] = true
		if n == 1 {
			return true
		}
	}
}

func main() {
	fmt.Printf("%v", isHappy(123456))
}
