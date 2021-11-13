package main

import "fmt"

//Abs - return absolute value of input int
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsNearbyDuplicate(nums []int, k int) bool {
	myMap := make(map[int]int)
	for idx, val := range nums {
		if _, ok := myMap[val]; ok && idx != myMap[val] && Abs(myMap[val]-idx) <= k {
			return true
		}
		myMap[val] = idx
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
