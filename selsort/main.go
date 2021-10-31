package main

import (
	"fmt"

	"github.com/walkrist/goalgs/structgenerator"
)

func selectionSort(unorderedList []int) []int {
	k := 0
	len := len(unorderedList)
	fmt.Printf("Initial list is %v\n\n", unorderedList)
	for j := k; j < len; j++ {
		min := unorderedList[k]
		fmt.Printf("Starting %v iteration, min is on %v index it has value of %v\n\n", j+1, k, min)
		for idx, val := range unorderedList[k+1:] {
			if val <= min {
				fmt.Printf("Found %v at index %v less or equal than %v, switching their positions...\n\n", val, idx+k+1, min)
				min = val
				unorderedList[k+1:][idx] = unorderedList[k]
				unorderedList[k] = min
				fmt.Printf("List after modifications: %v\n\n", unorderedList)
			}
		}
		k++
	}
	orderedList := unorderedList
	fmt.Printf("List has been sorted, result is: %v\n\n", orderedList)
	return orderedList
}

func main() {
	list := structgenerator.UnorderedIntList(15)
	selectionSort(list)
}
