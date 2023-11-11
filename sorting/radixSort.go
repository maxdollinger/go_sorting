package sorting

import (
	"time"
)

const BUCKET_SIZE = 10
const UNIX_DIGITS = 1000000000

func RadixSort(s []time.Time) []time.Time {

	for d := 1; d <= UNIX_DIGITS; d *= BUCKET_SIZE {

		C := [BUCKET_SIZE]uint{}
		for i := range s {
			C[extractKey(s[i], d)]++
		}

		for i := 1; i < BUCKET_SIZE; i++ {
			C[i] += C[i-1]
		}

		tmp := make([]time.Time, len(s))
		for i := (len(s) - 1); i >= 0; i-- {
			k := extractKey(s[i], d)
			C[k]--
			tmp[C[k]] = s[i]
		}

		s = tmp
	}

	return s
}

func extractKey(t time.Time, d int) int {
	return int(t.Unix()) / d % BUCKET_SIZE
}
