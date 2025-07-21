package sorting

// WorkItem represents a slice segment to be processed
type WorkItem[T Sortable] struct {
	bucket []T
	digits int64 // current digit position
}

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

	// Initialize stack with the entire slice
	stack := []WorkItem[T]{{bucket: s[:], digits: digits}}

	// Process stack until empty
	for len(stack) > 0 {
		// Pop work item from stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		afsWorkerIterative(item.bucket, item.digits, &stack)
	}
}

func afsWorkerIterative[S Sortable](s []S, d int64, stack *[]WorkItem[S]) {
	if len(s) <= 1 || d == 0 {
		return
	}

	// Count frequency of each digit
	C := [AFS_RADIX]int{}
	for i := range s {
		key := afsKey(s[i], d)
		C[key]++
	}

	// Calculate start and end indices for each bucket (relative to subslice)
	H := [AFS_RADIX]int{0}
	T := [AFS_RADIX]int{C[0]}
	for i := 1; i < AFS_RADIX; i++ {
		H[i] = T[i-1]
		T[i] = H[i] + C[i]
	}

	// Partition elements into buckets using in-place swapping
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
		for i := len(C) - 1; i >= 0; i-- {
			if C[i] > 1 {
				*stack = append(*stack, WorkItem[S]{bucket: s[(T[i] - C[i]):T[i]], digits: d})
			}
		}
	}
}

func afsKey(t Sortable, d int64) int {
	if d == 0 {
		return 0
	}

	return int((t.SortValue() / d) % AFS_RADIX)
}
