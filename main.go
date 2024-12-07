package main

import "fmt"

func twoSum(nums []int, target int) []int {
	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr = append(arr, nums[i])
				arr = append(arr, nums[j])
			}
		}

	}
	return arr
}

func main() {
	fmt.Printf("%T\n", twoSum([]int{1, 2, 3}, 4))
}
