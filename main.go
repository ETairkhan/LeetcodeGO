package main

import (
	"fmt"
	"golang/leetcodeProblems"
)

func main(){
	nums := []int{1,3,5,6}
	target := 1
	fmt.Println(leetcodeProblems.SearchInsert(nums,target))
}