package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	var resultSet []string
	var minIdxSum int
	map2 := make(map[string]int)
	for idx, rest := range list2 {
		map2[rest] = idx
	}
	for idx, rest := range list1 {
		if val, ok := map2[rest]; ok {
			if resultSet == nil {
				minIdxSum = idx + val
			}
			fmt.Printf("Found %v\n\n", rest)
			if idx+val == minIdxSum {
				resultSet = append(resultSet, rest)
				fmt.Printf("minidx equal %v\n\n", resultSet)
			} else if idx+val < minIdxSum {
				minIdxSum = idx + val
				resultSet = nil
				resultSet = append(resultSet, rest)
				fmt.Printf("newminidx %v\n\n", resultSet)
			}
		}
	}
	return resultSet
}

func main() {
	fmt.Print(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KNN", "KFC", "Burger King", "Tapioca Express", "Shogun"}))
}
