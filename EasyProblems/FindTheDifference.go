package EasyProblems

func FindTheDifference(s string, t string) byte {
	charCount := map[byte]int{}

	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		charCount[t[i]]--
		if charCount[t[i]] < 0 {
			return t[i]
		}
	}
	return 0
}
