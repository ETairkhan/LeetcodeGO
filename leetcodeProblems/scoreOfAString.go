package leetcodeProblems

func ScoreOfString(s string) int {
	result := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			result += int(s[i] - s[i+1])
		} else {
			result += int(s[i+1] - s[i])
		}

	}
	return result
}
