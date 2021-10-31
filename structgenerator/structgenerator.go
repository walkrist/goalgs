package structgenerator

import (
	"fmt"
	"math/rand"
)

// OrderedIntList - generates simple ordered list of integers
// from start integer with given lenght and step, if step == 0 then
// random step will be used.
// Random step - random int in range [0,10) will be added to previous value
func OrderedIntList(start int, len int, step int) []int {
	nums := make([]int, len)
	switch {
	case step < 0:
		fmt.Printf("Invalid step = %v, please provide non-negative step", step)
	case step == 0:
		for i := 0; i < len; i++ {
			nums[i] = start
			start = start + rand.Intn(10)
		}
	default:
		for i := 0; i < len; i++ {
			nums[i] = start
			start = start + step
		}
	}
	return nums
}

// UnorderedIntList - generates list of given length
// containing random integers in range [0,100)
func UnorderedIntList(len int) []int {
	nums := make([]int, len)
	for i := 0; i < len; i++ {
		nums[i] = rand.Intn(100)
	}
	return nums
}
