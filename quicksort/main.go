package main

import (
	"fmt"
	"log"

	"github.com/walkrist/goalgs/structgenerator"
)

//Hoare Partitioning

func hoarePartitioning(U []int, low, hi int) int {
	i, j := low, hi
	pivot := U[(j+i)/2]
	fmt.Printf("Partitioning of %v\n", U)
	fmt.Printf("Pivot element is U[%v] = %v\n\n", (j+i)/2, pivot)
	for {
		for U[i] < pivot {
			i++
		}
		for U[j] > pivot {
			j--
		}
		if i >= j {
			fmt.Printf("Result: %v\n\n", U)
			return j
		}
		U[i], U[j] = U[j], U[i]
		i++
		j--
	}
}

func lomutoPartitioning(U []int, low, hi int) int {
	i := low
	pivot := U[hi]

	for j := low; j < hi; j++ {
		if U[j] < pivot {
			U[i], U[j] = U[j], U[i]
			i++
		}
	}
	U[hi], U[i] = U[i], U[hi]
	return i
}

func quicksort(U []int, low, hi int, sortType string) {
	if low < hi {
		if sortType == "Lomuto" || sortType == "lomuto" {
			p := lomutoPartitioning(U, low, hi)
			fmt.Printf("Sorting left subarray: %v\n\n", U)
			quicksort(U, low, p-1, sortType)
			fmt.Printf("Sorting right subarray: %v\n\n", U)
			quicksort(U, p+1, hi, sortType)

		} else if sortType == "Hoare" || sortType == "hoare" {
			p := hoarePartitioning(U, low, hi)
			fmt.Printf("Sorting left subarray: %v\n\n", U)
			quicksort(U, low, p, sortType)
			fmt.Printf("Sorting right subarray: %v\n\n", U)
			quicksort(U, p+1, hi, sortType)
		} else {
			err := "Invalid sort type"
			log.Fatalf("%v: %v", err, sortType)
		}

	}
}

func main() {
	l := 10
	lst := structgenerator.UnorderedIntList(l)
	quicksort(lst, 0, l-1, "Hoare")
	fmt.Printf("Hoare partitioning\n--------\nSorted list:%v\n\n", lst)
	quicksort(lst, 0, l-1, "Lomuto")
	fmt.Printf("Lomuto partitioning\n--------\nSorted list:%v\n\n", lst)
}
