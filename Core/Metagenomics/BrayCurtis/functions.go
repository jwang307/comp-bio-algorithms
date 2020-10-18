package BrayCurtis

// Write the function BrayCurtisDistance() with the following input and output
// Input: two frequency maps sample1 and sample2
// Output: the Bray-Curtis distance between sample1 and sample2

func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	minSum := 0
	map1Sum := 0
	map2Sum := 0

	for val, _ := range sample1 {
		map1Sum += sample1[val]
		if sample1[val] < sample2[val] {
			minSum += sample1[val]
		} else {
			minSum += sample2[val]
		}
	}

	for val, _ := range sample2 {
		map2Sum += sample2[val]
	}

	return 1.0 - float64(minSum)/((float64(map1Sum)+float64(map2Sum))/2)
}
