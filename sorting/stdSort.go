package sorting

import (
	"sort"
)

func StandartSort[S Sortable](s []S) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].SortValue() < s[j].SortValue()
	})
}
