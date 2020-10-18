package Alignment

//EditDistance takes two strings as input. It returns the Levenshtein distance
//between the two strings; that is, the minimum number of substitutions, insertions, and deletions
//needed to transform one string into the other.

func EditDistance(str1, str2 string) int {
	//make initial 2 d array
	LCSMatrix := make([][]int, len(str1)+1)
	for i := range LCSMatrix {
		LCSMatrix[i] = make([]int, len(str2)+1)
	}

	for i := range LCSMatrix {
		LCSMatrix[i][0] = i
	}

	for i := range LCSMatrix[0] {
		LCSMatrix[0][i] = i
	}

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				LCSMatrix[i][j] = LCSMatrix[i-1][j-1]
			} else {
				LCSMatrix[i][j] = 1 + Min(LCSMatrix[i][j-1], LCSMatrix[i-1][j], LCSMatrix[i-1][j-1])
			}
		}
	}
	return LCSMatrix[len(str1)][len(str2)]
}

func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("no inputs to Max")
	}
	m := 0
	for i, val := range nums {
		if val < m || i == 0 {
			m = val
		}
	}
	return m
}
