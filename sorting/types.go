package sorting

type Sortable interface {
	SortValue() int64
}

type SortingFn func(s []Sortable)
