package leetcodeProblems

import "fmt"

func Rotate(nums []int, k int) {
	answer := []int{}
	for i := 0; i < k; i++ {
		answer = append(answer, nums[len(nums)-i])
		answer = append(answer ,nums[:len(nums)-i]...)
		fmt.Println(answer )
	}
}
