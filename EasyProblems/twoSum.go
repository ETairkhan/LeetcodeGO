package EasyProblems

func TwoSum(nums []int, target int) []int {
	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr[0] = nums[i]
				arr[1] = nums[j]
			}
		}

	}
	return arr
}

//func main() {
//	fmt.Printf("%T\n", twoSum([]int{1, 2, 3}, 4))
//}
