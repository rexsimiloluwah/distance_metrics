package text

// Computes the Levensthein distance between two strings
// The minimum number of single character edits to convert one string to another
// edit operations can be (insertion,deletion, or substition)
// Time complexity : O(m*n), where m and n are the lengths of string 1 and 2
// Space complexity : O(m*n), a new array is created.
func Levensthein(s1 string, s2 string) int {
	// Instantiate the Levensthein matrix
	if len(s1) == 0 {
		return len(s2)
	}

	if len(s2) == 0 {
		return len(s1)
	}

	levArr := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		levArr[i] = make([]int, len(s2)+1)
		for j := 0; j < len(s2)+1; j++ {
			levArr[i][j] = 0
		}
	}

	for i := 0; i < len(s1)+1; i++ {
		levArr[i][0] = i
	}

	for i := 0; i < len(s2)+1; i++ {
		levArr[0][i] = i
	}

	// Fill in the character edit count using the levensthein algorithm
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				levArr[i][j] = levArr[i-1][j-1]
			} else {
				levArr[i][j] = MinVar(
					levArr[i][j-1],
					levArr[i-1][j-1],
					levArr[i-1][j],
				) + 1
			}
		}
	}

	// fmt.Println(levArr)
	return levArr[len(s1)][len(s2)]
}

//@utility : Find the minimum between a variadic number of inputs
func MinVar(values ...int) int {
	v := values[0]
	for _, val := range values {
		if val < v {
			v = val
		}
	}
	return v
}
