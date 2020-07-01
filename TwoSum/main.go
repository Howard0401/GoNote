package main

import "fmt"

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if target-nums[i] == nums[j] {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, ok := m[temp]; ok {
			return []int{m[temp], i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	fmt.Printf("%#v", twoSum(nums, 6))
}
