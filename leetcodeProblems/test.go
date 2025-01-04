package leetcodeProblems

import (
	"strconv"
	"strings"
)

func ConvertDateToBinary(date string) string {
	dict := strings.Split(date, "-") // Split the date into components
	answer := ""

	for i := 0; i < len(dict); i++ {
		num, err := strconv.Atoi(dict[i]) // Convert string to integer
		if err != nil {
			return "Error: Invalid input" // Handle error if conversion fails
		}
		binary := strconv.FormatInt(int64(num), 2) // Convert integer to binary
		if i > 0 {
			answer += "-" // Add a separator between binary components
		}
		answer += binary
	}

	return answer
}
