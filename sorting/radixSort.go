package sorting

const BUCKET_SIZE = 1000

func RadixSort[T Sortable](s []T) {
	if len(s) == 0 {
		return
	}

	maxVal := getMaxValue(s)
	tmp := make([]T, len(s))

	for exp := int64(1); maxVal/exp > 0; exp *= BUCKET_SIZE {

		C := [BUCKET_SIZE]int{}
		for i := range s {
			digit := extractKey(s[i], exp)
			C[digit]++
		}

		for i := 1; i < BUCKET_SIZE; i++ {
			C[i] += C[i-1]
		}

		for i := len(s) - 1; i >= 0; i-- {
			digit := extractKey(s[i], exp)
			C[digit]--
			tmp[C[digit]] = s[i]
		}

		copy(s, tmp)
	}
}

func extractKey(t Sortable, exp int64) int {
	return int((t.SortValue() / exp) % BUCKET_SIZE)
}

func getMaxValue[T Sortable](s []T) int64 {
	if len(s) == 0 {
		return 0
	}

	max := s[0].SortValue()
	for i := 1; i < len(s); i++ {
		if val := s[i].SortValue(); val > max {
			max = val
		}
	}

	return max
}
