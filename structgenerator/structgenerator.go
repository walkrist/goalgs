package structgenerator

//OrderedIntList - generates simple ordered list of integers in range
// from fromNum to toNum.
func OrderedIntList(fromNum int, toNum int) []int {
	len := (toNum - fromNum) + 1
	nums := make([]int, len)
	for i := 0; i < len; i++ {
		nums[i] = fromNum
		fromNum++
	}
	return nums
}
