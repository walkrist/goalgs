package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var resultArray []int
	myMap := make(map[int]int)
	ptr1 := &nums1
	ptr2 := &nums2
	if len(nums1) > len(nums2) {
		ptr1 = &nums2
		ptr2 = &nums1
	}
	for _, val := range *ptr1 {
		myMap[val]++
	}
	for _, val := range *ptr2 {
		if myMap[val] > 0 {
			resultArray = append(resultArray, val)
			myMap[val]--
		}
	}
	return resultArray
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
