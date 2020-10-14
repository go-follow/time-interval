package time_interval

import "time"

type SortType int

const (
	// Ascending sort ascending (default)
	Ascending SortType = iota
	// Descending sort descending
	Descending
)

func beforeOrEqual(t1, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}

func afterOrEqual(t1, t2 time.Time) bool {
	return t1.After(t2) || t1.Equal(t2)
}
