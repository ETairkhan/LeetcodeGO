package leetcodeProblems
func FinalValueAfterOperations(operations []string) int {
  answer := 0
  for i := 0; i <len(operations); i++{
      if operations[i] == "++X" || operations[i] == "X++"{
          answer++
      }else if operations[i] =="--X" || operations[i] =="X--"{
          answer--
      }
  }
  return answer
}