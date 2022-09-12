package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	target := 6
	indexes := twoSum(nums, target)
	fmt.Printf("solution is %v\n", indexes)
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		if p, ok := hashTable[target-v]; ok {
			return []int{p, i}
		}
		hashTable[v] = i
	}
	return nil
}
