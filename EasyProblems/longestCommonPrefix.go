package EasyProblems

func LongestCommonPrefix(strs []string) string {
	answer := ""

	for i := 0; i < len(strs[0]); i++ {
		answer += string(strs[0][i])
	}
	for i := 1; i < len(strs); i++ {
		if len(answer) > len(strs[i]) {
			answer = answer[:len(strs[i])]
		}
		for j := 0; j < len(answer); j++ {
			if answer[j] != strs[i][j] {
				if j == 0 {
					return ""
				}
				answer = answer[:j]
			}
		}
	}
	return answer
}
