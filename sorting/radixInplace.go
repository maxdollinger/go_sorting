package sorting

import (
	"sort"
	"sync"
	"time"
)

func RadixSortInplace(s []time.Time) []time.Time {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	radixInplace(s, 1000000000, wg)
	wg.Wait()

	return s
}

func radixInplace(s []time.Time, faktor int, wg *sync.WaitGroup) {
	defer wg.Done()

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
				if C[i] > 100 {
					wg.Add(1)
					go radixInplace(s[(T[i]-C[i]):T[i]], faktor, wg)
				} else {
					radixStd(s[(T[i] - C[i]):T[i]])
				}
			}
		}
	}
}

func radixStd(s []time.Time) {
	sort.Slice(s, func(i, j int) bool { return s[i].Unix() < s[j].Unix() })
}
