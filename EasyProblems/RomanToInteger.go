package EasyProblems

func RomanToInt(s string) int {
	myMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	answer := 0
	for i := 0; i < len(s); i++ {
		value1 := myMap[string(s[i])]
		if i < len(s)-1 {
			value2 := myMap[string(s[i+1])]
			if value2 > value1 {
				answer -= value1
				continue
			}
		}
		answer += value1
	}
	return answer

}
