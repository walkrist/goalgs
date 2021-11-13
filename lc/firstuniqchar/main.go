package main

import "fmt"

func firstUniqChar(s string) int {
	uniqIdx := -1
	myMap := make(map[string]int)
	for _, val := range s {
		myMap[string(val)]++
	}
	for pos, val := range s {
		if myMap[string(val)] == 1 {
			uniqIdx = pos
			return uniqIdx
		}
	}
	return uniqIdx
}

func main() {
	fmt.Print(firstUniqChar("aaabb"))
}
