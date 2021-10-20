package main

import (
	"fmt"
)

func nums_set(fromNum int, toNum int) []int {
	len := (toNum - fromNum) + 1
	nums := make([]int, len)
	for i := 0; i < len; i++ {
		nums[i] = fromNum
		fromNum++
	}
	return nums
}

func num_search(numSet []int, numFind int) bool {
	suggest := len(numSet) / 2
	tries := 0
	fmt.Printf("Initial set is %v.\nSuggestion is %v index\n", numSet, suggest)
	for {
		tries++
		switch {
		case numSet[suggest] == numFind:
			fmt.Printf("Success on %v try. Found it at %v index.\nSet was %v\n", tries, suggest, numSet)
			return true
		case numSet[suggest] != numFind:
			fmt.Printf("Failed on %v try. It's not on %v index.\nSet was %v\n", tries, suggest, numSet)
			switch {
			case len(numSet) == 1:
				fmt.Printf("No luck after %v tries. Its not here\n", tries)
				return false
			case numSet[suggest] > numFind:
				numSet = numSet[:suggest+1]
			case numSet[suggest] < numFind:
				numSet = numSet[suggest:]
			}
			suggest = len(numSet) / 2
		}
	}
}

func main() {
	set := nums_set(0, 1024)
	num_search(set, 402)
}
