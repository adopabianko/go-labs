package main

import "fmt"

func main() {
	result := twoSum([]int{3, 2, 3}, 6)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, v := range nums {
		n := target - v

		if _, ok := hashmap[n]; ok {
			return []int{hashmap[n], i}
		}

		hashmap[v] = i
	}

	return []int{}
}
