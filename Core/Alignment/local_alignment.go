package Alignment

//LocalAlignment takes two strings, along with match, mismatch, and gap scores.
//It returns a maximum score local alignment of the strings corresponding to these penalties.

func LocalAlignment(str0, str1 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
	LocalScoreMatrix := LocalScoreTable(str0, str1, match, mismatch, gap)

	MaxScore := 0.0

	EndI := 0
	EndJ := 0

	for i := 0; i <= len(str0); i++ {
		for j := 0; j <= len(str1); j++ {
			if LocalScoreMatrix[i][j] > MaxScore {
				MaxScore = LocalScoreMatrix[i][j]
				EndI = i
				EndJ = j
			}
		}
	}

	backtrack := LocalBacktrack(str0, str1, match, mismatch, gap)

	return OutputLocalAlignment(str0, str1, backtrack, EndI, EndJ)

}

func OutputLocalAlignment(str0, str1 string, backtrack [][]string, i, j int) (Alignment, int, int, int, int) {
	var a Alignment

	EndI := i
	EndJ := j

	//we want a[0] to be the top row and a[1] to be bottom row of alignment
	for (backtrack[i][j] != "ZERO") && (i > 0) && (j > 0) {
		//chop off symbols of two strings as we go
		//start at bottom right of our current backtrack table
		if backtrack[i][j] == "UP" {
			//a[0] // first string in Alignment
			//a[1] // second string in the Alignment
			a[0] = string(str0[i-1]) + a[0]
			a[1] = "-" + a[1]
			i--

		} else if backtrack[i][j] == "LEFT" {
			a[1] = string(str1[j-1]) + a[1]
			a[0] = "-" + a[0]
			j--
		} else { // "DIAG"
			a[0] = string(str0[i-1]) + a[0]
			a[1] = string(str1[j-1]) + a[1]
			i--
			j--
		}
	}
	return a, i, EndI, j, EndJ
}

func LocalBacktrack(str0, str1 string, match, mismatch, gap float64) [][]string {
	if len(str0) == 0 || len(str1) == 0 {
		panic("Blah")
	}
	numRows := len(str0) + 1
	numCols := len(str1) + 1

	backtrack := make([][]string, numRows)
	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}

	// grab the scores in the alignment table
	scoreTable := LocalScoreTable(str0, str1, match, mismatch, gap)
	// 2-D array of float64

	//first, let's set the backtracking pointers of 0-th row and column
	for j := 0; j < numCols; j++ {
		backtrack[0][j] = "ZERO"
	}
	for i := 0; i < numRows; i++ {
		backtrack[i][0] = "ZERO"
	}

	// traverse rest of scoreTable, and figure out which previous value was used to update each value
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			// check what value was used
			if scoreTable[i][j] == scoreTable[i-1][j]-gap {
				// indel
				backtrack[i][j] = "UP"
			} else if scoreTable[i][j] == scoreTable[i][j-1]-gap {
				backtrack[i][j] = "LEFT"
			} else if scoreTable[i][j] == 0 {
				backtrack[i][j] = "ZERO"
			} else {
				backtrack[i][j] = "DIAG"
			}
		}
	}
	return backtrack
}
