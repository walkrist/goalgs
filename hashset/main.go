package main

import "fmt"

// Simpliest hashSet - it will store only integers

// MyHashSet - a structure to keep set's data
type MyHashSet struct {
	data []bool
}

//Constructor - way to initialize the set
func Constructor() MyHashSet {
	setLen := 1000001
	return MyHashSet{
		data: make([]bool, setLen),
	}
}

// Add - adds the key to set
func (set *MyHashSet) Add(key int) {
	set.data[key] = true
}

// Remove -removes the key from set
func (set *MyHashSet) Remove(key int) {
	set.data[key] = false
}

// Contains - check if set contains an element
func (set *MyHashSet) Contains(key int) bool {
	return set.data[key]
}

// Keys - list all keys that currently in set
func (set *MyHashSet) Keys() []int {
	var resultingSet []int
	for i := 0; i < len(set.data); i++ {
		if set.data[i] {
			resultingSet = append(resultingSet, i)
		}
	}
	return resultingSet
}

func main() {
	obj := Constructor()
	fmt.Printf("%v\n\n", obj.Keys())
	obj.Add(1)
	fmt.Printf("%v\n\n", obj.Keys())
	obj.Add(2)
	fmt.Printf("%v\n\n", obj.Keys())
	obj.Remove(2)
	fmt.Printf("%v\n\n", obj.Keys())
	println(obj.Contains(2))
	println(obj.Contains(1))
}
