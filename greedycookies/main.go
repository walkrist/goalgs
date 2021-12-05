// Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie.

// Each child i has a greed factor g[i], which is the minimum size of a cookie that the child will be content with; and each cookie j has a size s[j]. If s[j] >= g[i], we can assign the cookie j to the child i, and the child i will be content. Your goal is to maximize the number of your content children and output the maximum number.

// Input: g = [1,2,3], s = [1,1]
// Output: 1
// Explanation: You have 3 children and 2 cookies. The greed factors of 3 children are 1, 2, 3.
// And even though you have 2 cookies, since their size is both 1, you could only make the child whose greed factor is 1 content.
// You need to output 1

package main

import (
	"fmt"
	"sort"
)

// my first solution <not optimal>
func findContentChildren(g []int, s []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(g)))
	i := 0
	count := 0
	for i < len(g) {
		for idx, cookie := range s {
			if cookie >= g[i] {
				s[idx] = 0
				count++
				break
			}
		}
		i++
	}
	return count
}

//optimal solution

func findContentChildrenOptimal(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si, res := 0, 0, 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			res++
			si++
			gi++
		} else {
			si++
		}
	}
	return res
}

func main() {
	fmt.Printf("%v", findContentChildren([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 1}))
}
