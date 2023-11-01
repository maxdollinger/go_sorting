package sorting

import (
	"sort"
	"time"
)

func RadixSortInplace(s []time.Time) []time.Time {

	radixInplace(s, 10000000000)

	return s
}

func radixInplace(s []time.Time, faktor int) {

	if sort.SliceIsSorted(s, func(i, j int) bool { return s[i].Unix() < s[j].Unix() }) {
		return
	}

	C := [BUCKET_SIZE]int{}
	for i := range s {
		C[extractKey(s[i], faktor)]++
	}

	H := [BUCKET_SIZE]int{0}
	T := [BUCKET_SIZE]int{C[0]}
	for i := 1; i < BUCKET_SIZE; i++ {
		H[i] = T[i-1]
		T[i] = H[i] + C[i]
	}

	for i := range s {
		k := extractKey(s[i], faktor)
		for H[k] < T[k] && i != H[k] {
			s[i], s[H[k]] = s[H[k]], s[i]
			H[k]++
			k = extractKey(s[i], faktor)
		}
	}

	faktor = faktor / BUCKET_SIZE
	if faktor > 0 {
		for i := range C {
			if C[i] > 0 {
				if C[i] > 1000 {
					radixInplace(s[(T[i]-C[i]):T[i]], faktor)
				} else {
					StandartSort(s[(T[i] - C[i]):T[i]])
				}
			}
		}
	}

}
