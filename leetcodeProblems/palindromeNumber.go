package leetcodeProblems

import (
	"fmt"
	"strconv"
)

func IsPalindrome(x int) bool {
	arr := strconv.Itoa(x)
	var answer []rune
	for i := len(arr) - 1; i >= 0; i-- {
		answer = append(answer, rune(arr[i]))
	}
	fmt.Println(answer)
	for i := len(arr) - 1; i >= 0; i--{
		if answer[i] != rune(arr[i]){
			return false
		}
	}
	return true
}
