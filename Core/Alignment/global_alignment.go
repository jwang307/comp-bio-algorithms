package Alignment

import "fmt"

type Alignment [2]string

//GlobalAlignment takes two strings, along with match, mismatch, and gap scores.
//It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str0, str1 string, match, mismatch, gap float64) Alignment {
	backtrack := GlobalBacktrack(str0, str1, match, mismatch, gap)
	// backtrack will compute scores using recurrences and give me all the backtracking pointers as a 2-D array.
	optAlignment := OutputGlobalAlignment(str0, str1, backtrack)
	// this function will use the backtracking info to construct our alignment
	return optAlignment
}

//OutputGlobalAlignment takes two strings and a backtrack matrix of strings. It returns the (presumably optimal) alignment corresponding to this backtracking.
func OutputGlobalAlignment(str0, str1 string, backtrack [][]string) Alignment {
	var a Alignment

	//we want a[0] to be the top row and a[1] to be bottom row of alignment
	for len(str0) > 0 || len(str1) > 0 {
		//chop off symbols of two strings as we go
		//start at bottom right of our current backtrack table
		rowIndex := len(str0)
		colIndex := len(str1)
		if backtrack[rowIndex][colIndex] == "UP" {
			//a[0] // first string in Alignment
			//a[1] // second string in the Alignment
			a[0] = string(str0[rowIndex-1]) + a[0]
			a[1] = "-" + a[1]
			str0 = str0[:len(str0)-1]

		} else if backtrack[rowIndex][colIndex] == "LEFT" {
			a[1] = string(str1[colIndex-1]) + a[1]
			a[0] = "-" + a[0]
			str1 = str1[:len(str1)-1]
		} else { // "DIAG"
			a[0] = string(str0[rowIndex-1]) + a[0]
			a[1] = string(str1[colIndex-1]) + a[1]
			str0 = str0[:len(str0)-1]
			str1 = str1[:len(str1)-1]
		}
	}
	return a
}

//GlobalBacktrack takes two strings and alignment penalties/rewards.
//It returns a 2-D slice of strings giving backtracking pointers for an optimal
//global alignment of the two strings using these penalties/rewards.
func GlobalBacktrack(str0, str1 string, match, mismatch, gap float64) [][]string {
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
	scoreTable := GlobalScoreTable(str0, str1, match, mismatch, gap)
	// 2-D array of float64

	//first, let's set the backtracking pointers of 0-th row and column
	for j := 1; j < numCols; j++ {
		backtrack[0][j] = "LEFT"
	}
	for i := 1; i < numRows; i++ {
		backtrack[i][0] = "UP"
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
			} else {
				// I should be more precise than this
				backtrack[i][j] = "DIAG"
			}
		}
	}
	return backtrack
}

//PrintAlignment takes an alignment and prints the two corresponding rows of the alignment.
func PrintAlignment(a Alignment) {
	fmt.Println(a[0])
	fmt.Println(a[1])
}
