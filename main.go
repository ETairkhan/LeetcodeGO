package main

import (
	"fmt"
	"golang/leetcodeProblems"
)

func main() {
	var nums []int
	nums = append(nums, 1, 1, 2)
	fmt.Println(nums)

	k := leetcodeProblems.RemoveDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(k)
}
