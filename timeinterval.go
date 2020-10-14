package time_interval

import "time"

type SortType int

const (
	// Increase sort by increase (default)
	Increase SortType = iota
	// Decrease sort by decrease
	Decrease
)

func beforeOrEqual(t1, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}

func afterOrEqual(t1, t2 time.Time) bool {
	return t1.After(t2) || t1.Equal(t2)
}
