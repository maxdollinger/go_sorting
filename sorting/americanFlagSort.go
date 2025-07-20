package sorting

// sorts each subslice of s in a separate goroutine for needs significantly more memory for larger slices
func AmericanFlagSort[T Sortable](s []T) {
	if len(s) == 0 {
		return
	}

	maxVal := getMaxValue(s)
	if maxVal == 0 {
		return
	}

	// Find the highest digit position to start from
	digits := int64(1)
	for maxVal/digits >= int64(AFS_RADIX) {
		digits *= int64(AFS_RADIX)
	}

	afsWorker(s, digits)
}

func afsWorker[S Sortable](s []S, d int64) {
	if len(s) <= 1 || d == 0 {
		return
	}

	// Count frequency of each digit
	C := [AFS_RADIX]int{}
	for i := range s {
		key := afsKey(s[i], d)
		C[key]++
	}

	// Calcutlate start and end indices for each bucket
	H := [AFS_RADIX]int{0}
	T := [AFS_RADIX]int{C[0]}
	for i := 1; i < AFS_RADIX; i++ {
		H[i] = T[i-1]
		T[i] = H[i] + C[i]
	}
	// Partition elements into buckets using in-place swapping
	for i := range s {
		// k Bucket key and; H[k] start idx of bucket; T[k] end idx of bucket
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
			afsWorker(s[(T[i]-c):T[i]], d)
		}
	}
}

func afsKey(t Sortable, d int64) int {
	if d == 0 {
		return 0
	}

	return int((t.SortValue() / d) % AFS_RADIX)
}
