package sorting

import (
	"sync"
	"time"
)

const AFS_BUCKET_SIZE = 500
const CUTOFF = 25
const UNIX_DIGITS = 1000000000

func AmericanFlagSort(s []time.Time) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go radixRec(s, UNIX_DIGITS, wg)
	wg.Wait()
}

func radixRec(s []time.Time, d int, wg *sync.WaitGroup) {
	defer wg.Done()

	C := [AFS_BUCKET_SIZE]int{}
	for i := range s {
		C[afsKey(s[i], d)]++
	}

	H := [AFS_BUCKET_SIZE]int{0}
	T := [AFS_BUCKET_SIZE]int{C[0]}
	for i := 1; i < AFS_BUCKET_SIZE; i++ {
		H[i] = T[i-1]
		T[i] = H[i] + C[i]
	}

	for i := range s {
		k := afsKey(s[i], d)
		for H[k] < T[k] && i != H[k] {
			s[i], s[H[k]] = s[H[k]], s[i]
			H[k]++
			k = afsKey(s[i], d)
		}
	}

	d = d / AFS_BUCKET_SIZE
	if d > 0 {
		for i, c := range C {
			if c > CUTOFF {
				wg.Add(1)
				go radixRec(s[(T[i]-c):T[i]], d, wg)
			} else if c > 1 {
				InsertionSort(s[(T[i] - c):T[i]])
			}
		}
	}

}

func afsKey(t time.Time, d int) int {
	return int(t.Unix()) / d % AFS_BUCKET_SIZE
}
