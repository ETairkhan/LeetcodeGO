package leetcodeProblems

func ContainsDuplicate(nums []int) bool {

	for i:= 0; i < len(nums)-1; i++{
		for j:= i; j < len(nums) - 1; j++{
			if nums[i] == nums[j + 1]{
				return true
			}
		}
	}
	return false
}
