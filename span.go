package time_interval

import "time"

// Span time interval
type Span struct {
	start time.Time
	end   time.Time
}

// New initialization of a new time interval
// the beginning of the interval must necessarily be less than the end of the interval
func New(start, end time.Time) Span {
	if afterOrEqual(start, end) {
		panic("time start cannot be more time end")
	}
	return Span{
		start: start,
		end:   end,
	}
}

// Equal full equals of two time slots
func (s *Span) Equal(input Span) bool {
	return s.start.Equal(input.start) && s.end.Equal(input.end)
}

// IsIntersection - проверка на пересечение временных интервалов (не включаются случаи на стыке)
func (s *Span) IsIntersection(input Span) bool {
	return s.start.Before(input.end) && s.end.After(input.start)
}
