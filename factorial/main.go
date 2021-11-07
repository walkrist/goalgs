package main

func factorial(number int) int {
	if number == 1 {
		return number
	}
	return factorial(number-1) * (number)
}

func main() {
	print(factorial(10))
}
