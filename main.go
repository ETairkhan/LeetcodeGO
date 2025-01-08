package main

import (
	"fmt"
	"golang/leetcodeProblems"
)

func main() {
	operations := []string{"--X", "++X", "X++"}
	fmt.Println(leetcodeProblems.FinalValueAfterOperations(operations))
}
