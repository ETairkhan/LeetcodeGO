package EasyProblems

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Pointer for the position of the last unique element
	uniqueIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	// The count of unique elements is uniqueIndex + 1
	return uniqueIndex + 1
}
