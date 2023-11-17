package sorting

import (
	"time"
)

const BUCKET_SIZE = 10000

func RadixSort(s []time.Time) {

	tmp := make([]time.Time, len(s))
	for d := 1; d <= UNIX_DIGITS; d *= BUCKET_SIZE {

		C := [BUCKET_SIZE]int{}
		for i := range s {
			C[extractKey(s[i], d)]++
		}

		for i := 1; i < BUCKET_SIZE; i++ {
			C[i] += C[i-1]
		}

		for i := (len(s) - 1); i >= 0; i-- {
			k := extractKey(s[i], d)
			C[k]--
			tmp[C[k]] = s[i]
		}

		copy(s, tmp)
	}

}

func extractKey(t time.Time, d int) int {
	return int(t.Unix()) / d % BUCKET_SIZE
}
