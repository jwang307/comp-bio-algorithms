package Alignment

func LocalScoreTable(str0, str1 string, match, mismatch, gap float64) [][]float64 {

	numRows := len(str0) + 1
	numCols := len(str1) + 1

	scoringMatrix := make([][]float64, numRows)
	for i := range scoringMatrix {
		scoringMatrix[i] = make([]float64, numCols)
	}

	//first, penalize the 0-th row and column
	for j := range scoringMatrix {
		scoringMatrix[j][0] = 0
	}
	for j := range scoringMatrix[0] {
		scoringMatrix[0][j] = 0
	}

	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//need an up, left, and a diag
			up := scoringMatrix[i-1][j] - gap //indel
			left := scoringMatrix[i][j-1] - gap
			diag := scoringMatrix[i-1][j-1]
			if str0[i-1] == str1[j-1] {
				//match!
				diag += match
			} else {
				//mismatch!
				diag -= mismatch
			}
			// now take max
			scoringMatrix[i][j] = MaxFloat(up, left, diag, 0.0)
		}
	}

	return scoringMatrix

}
