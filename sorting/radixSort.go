package sorting

import "time"

const BUCKET_SIZE = 1000

func RadixSort(s []time.Time) []time.Time {

	n := len(s)
	faktor := 1
	for i := 0; i < 4; i++ {

		C := [BUCKET_SIZE]int{}
		for i := range s {
			C[extractKey(s[i], faktor)]++
		}

		for i := 1; i < BUCKET_SIZE; i++ {
			C[i] += C[i-1]
		}

		sorted := make([]time.Time, n)
		for i := (n - 1); i >= 0; i-- {
			k := extractKey(s[i], faktor)
			C[k]--
			sorted[C[k]] = s[i]
		}

		s = sorted
		faktor *= BUCKET_SIZE
	}

	return s
}

func extractKey(t time.Time, f int) int {
	return int(t.Unix()) / f % BUCKET_SIZE
}
