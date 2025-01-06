package leetcodeProblems


func Shuffle(nums []int, n int) []int {
	// Create a slice to store the result
result := make([]int, 0, 2*n)

// Iterate through the first half and second half of the array
for i := 0; i < n; i++ {
  result = append(result, nums[i])     // Append x_i
  result = append(result, nums[i+n])  // Append y_i
}

return result
}