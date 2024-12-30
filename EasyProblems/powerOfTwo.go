package EasyProblems

func IsPowerOfTwo(n int) bool {
	answer := 1
	for i:= 1; answer < n; i++{
		answer = answer * 2 
	}
	return answer == n
}
