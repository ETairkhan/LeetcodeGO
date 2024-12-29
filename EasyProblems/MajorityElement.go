package EasyProblems

func MajorityElement(nums []int) int {
	myMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]]++
	}
	biggest := 0
	answer := 0
	for num, count := range myMap {
		if count > biggest {
			biggest = count
			answer = num
		}
	}
	return answer
}
