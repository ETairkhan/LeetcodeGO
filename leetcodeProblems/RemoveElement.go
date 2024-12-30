package leetcodeProblems

func RemoveElement(nums []int, val int) int  {
	count := 0
	j := 0 
	for i := 0; i < len(nums); i++{
	  if nums[i] == val{
		  count++
	  
	  }else{
				
		  nums[j] = nums[i]
			j++
		}
	}
	 return (len(nums) - count)

}