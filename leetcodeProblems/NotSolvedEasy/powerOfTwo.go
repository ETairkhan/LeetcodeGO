package leetcodeProblems

func IsPowerOfTwo(n int) bool {
    if (n == 1){
		return true
	 }
	 if n % 2 == 0{
		return true
	 }else{
		return false
	 }
}