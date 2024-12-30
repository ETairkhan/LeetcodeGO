package EasyProblems

import "strconv"

func AddDigits(num int) int {
	str := strconv.Itoa(num)
	answer := 0
	for i := 0 ; i < len(str); i++{
		a , _ := strconv.Atoi(string(str[i]))
		answer += a
	}
	if answer >= 10 {
		answer = AddDigits(answer)
	}
	return answer
}
