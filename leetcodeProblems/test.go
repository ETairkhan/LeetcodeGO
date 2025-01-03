package leetcodeProblems

func StrStr(haystack string, needle string) int {
	if len(needle) == 0{
		return 0
	}
	
	for i := 0; i <= len(haystack) - len(needle); i++{
	 str1 := haystack[i:i+len(needle)]
	  if str1 == needle{
			return i
		 
	  }
  }
	return -1
}