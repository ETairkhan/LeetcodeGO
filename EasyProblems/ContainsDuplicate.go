package EasyProblems

func ContainsDuplicate(nums []int) bool {
	myMap := map[int]int{}
	for i:= 0; i < len(nums); i++{
		myMap[nums[i]]++
	}
	for i:= 0; i < len(nums); i++{
		if myMap[nums[i]] == 1{
			continue
		}else {
			return true
		}
	}
	return false
}
