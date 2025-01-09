package leetcodeProblems
func FindWordsContaining(words []string, x byte) []int {
    indices := []int{}
      for i, word := range words {
          for _, char := range word {
              if char == rune(x) {
                  indices = append(indices, i)
                  break
              }
          }
      }
      return indices
  }