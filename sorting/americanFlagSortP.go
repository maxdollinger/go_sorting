package sorting

import "sync"

// sorts each subslice of s in a separate goroutine for needs significantly more memory for larger slices
func AmericanFlagSortParallel[T Sortable](s []T) {
	if len(s) == 0 {
		return
	}

	maxVal := getMaxValue(s)
	if maxVal == 0 {
		return
	}

	digits := int64(1)
	for maxVal/digits >= int64(AFS_RADIX) {
		digits *= int64(AFS_RADIX)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go afsWorkerParallel(s, digits, &wg)
	wg.Wait()
}

func afsWorkerParallel[S Sortable](s []S, d int64, wg *sync.WaitGroup) {
	if len(s) <= 1 || d == 0 {
		return
	}

	C := [AFS_RADIX]int{}
	for i := range s {
		key := afsKey(s[i], d)
		C[key]++
	}

	H := [AFS_RADIX]int{0}
	T := [AFS_RADIX]int{C[0]}
	for i := 1; i < AFS_RADIX; i++ {
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

	d = d / AFS_RADIX
	if d > 0 {
		for i, c := range C {
			/* to reduce the number of goroutines, we only create a new goroutine if the bucket has more than 50 elements
			 * and use the fast insertion sort for buckets with less than 50 elements
			 */
			if c > 50 {
				wg.Add(1)
				go afsWorkerParallel(s[(T[i]-c):T[i]], d, wg)
			} else if c > 1 {
				InsertionSort(s[(T[i] - c):T[i]])
			}
		}
	}

	wg.Done()
}
