package sorting

func InsertionSort[S Sortable](s []S) {
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0 && s[j].SortValue() > s[j+1].SortValue(); j-- {
			s[j+1], s[j] = s[j], s[j+1]
		}
	}
}
