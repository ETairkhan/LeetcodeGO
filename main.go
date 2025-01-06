package main

import (
	"fmt"
	"golang/leetcodeProblems"
)

func main() {
	// Example 1
	nums1 := []int{2, 5, 1, 3, 4, 7}
	n1 := 3
	fmt.Println(leetcodeProblems.Shuffle(nums1, n1)) // Output: [2, 3, 5, 4, 1, 7]

	// Example 2
	nums2 := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n2 := 4
	fmt.Println(leetcodeProblems.Shuffle(nums2, n2)) // Output: [1, 4, 2, 3, 3, 2, 4, 1]

	// Example 3
	nums3 := []int{1, 1, 2, 2}
	n3 := 2
	fmt.Println(leetcodeProblems.Shuffle(nums3, n3)) // Output: [1, 2, 1, 2]
	
}
