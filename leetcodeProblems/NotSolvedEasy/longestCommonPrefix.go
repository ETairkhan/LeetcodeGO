package leetcodeProblems
// can solve this shit brooo
func LongestCommonPrefix(strs []string) string {
   answer := "" 
	for i:= 0; i < len(strs);i++{
		if len(strs[i])  < 3 || len(strs[i]) == 0 {
			return ""
		}
	}
	
	for i := 0; i< 3; i++{
	 	answer += string(strs[0][i]) 
	}

	for i := 1; i< len(strs); i++{
		for j := 0; j < 3; j++{
			if answer[j] != strs[i][j]{
				if j == 0 {
					return ""
				}
				answer = answer[:j]
			}
		}
		
	}
	 return answer
}