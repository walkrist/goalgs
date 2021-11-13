package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	var stringArr []string
	var sortedStr string
	for _, val := range s {
		stringArr = append(stringArr, string(val))
	}
	sort.Strings(stringArr)
	for _, let := range stringArr {
		sortedStr = sortedStr + let
	}
	return sortedStr
}

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)
	var resultArr [][]string
	for _, str := range strs {
		resultMap[sortString(str)] = append(resultMap[sortString(str)], str)
	}
	for _, v := range resultMap {
		resultArr = append(resultArr, v)
	}
	return resultArr
}

func main() {
	fmt.Println(groupAnagrams([]string{"a"}))
}
