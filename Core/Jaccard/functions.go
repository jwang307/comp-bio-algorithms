package Jaccard

// Write the function JaccardDistance() with the following input and output
// Input: two frequency maps sample1 and sample2
// Output: the Jaccard distance between sample1 and sample2

func JaccardDistance(sample1, sample2 map[string]int) float64 {
	minSum := 0.0
	maxSum := 0.0
	overlapString := make([]string, 0)

	for val, _ := range sample1 {
		if sample1[val] > sample2[val] {
			minSum += float64(sample2[val])
			maxSum += float64(sample1[val])
		} else {
			maxSum += float64(sample2[val])
			minSum += float64(sample1[val])
		}

		if sample2[val] > 0 {
			overlapString = append(overlapString, val)
		}
	}

	for val, _ := range sample2 {
		if !CheckOverlap(overlapString, val) {
			maxSum += float64(sample2[val])
		}
	}

	return 1.0 - minSum/maxSum
}

func CheckOverlap(reference []string, pattern string) bool {
	for _, str := range reference {
		if pattern == str {
			return true
		}
	}

	return false
}
