package sorting

import "time"

func InsertionSort(s []time.Time) {
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0 && s[j].Unix() > s[j+1].Unix(); j-- {
			s[j+1], s[j] = s[j], s[j+1]
		}
	}
}
