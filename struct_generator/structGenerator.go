package structGenerator

func orderedIntList(fromNum int, toNum int) []int {
	len := (toNum - fromNum) + 1
	nums := make([]int, len)
	for i := 0; i < len; i++ {
		nums[i] = fromNum
		fromNum++
	}
	return nums
}
