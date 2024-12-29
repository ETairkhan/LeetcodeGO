package EasyProblems

func IsValid(s string) bool {
	arr := make([]rune, len(s))
	index := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			arr[index] = rune(s[i])
			index++
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if index == 0 || !((arr[index-1] == '(' && s[i] == ')') || (arr[index-1] == '[' && s[i] == ']') || (arr[index-1] == '{' && s[i] == '}')) {
				return false
			}
			index--

		}
	}

	if index > 0 {
		return true
	}
	return true
}

//func main() {
//	fmt.Println(isValid("(]"))
//}
