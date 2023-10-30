package sorting

import (
	"sort"
	"time"
)

func StandartSort(s []time.Time) []time.Time {

	sort.Slice(s, func(i, j int) bool {
		return s[i].Unix() < s[j].Unix()
	})

	return s
}
