package time_interval

import "time"

func beforeOrEqual(t1, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}

func afterOrEqual(t1, t2 time.Time) bool {
	return t1.After(t2) || t1.Equal(t2)
}
