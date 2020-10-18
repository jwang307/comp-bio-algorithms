package Alignment

//CountSharedKmers takes two strings and an integer k. It returns the number of
//k-mers that are shared by the two strings.

func CountSharedKmers(str1, str2 string, k int) int {
	//create frequency maps
	freqMap1 := FrequencyMap(str1, k)
	freqMap2 := FrequencyMap(str2, k)

	sum := 0

	for i, _ := range freqMap1 {
		if freqMap1[i] < freqMap2[i] {
			sum += freqMap1[i]
		} else {
			sum += freqMap2[i]
		}
	}

	return sum
}
