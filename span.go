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

// Start - returning start time interval
func (s *Span) Start() time.Time {
	return s.start
}

// End - returning end time interval
func (s *Span) End() time.Time {
	return s.end
}

// IsEmpty - defines empty spacing
func (s Span) IsEmpty() bool {
	return s.start.IsZero() && s.end.IsZero()
}

// Equal full equals of two time slots
// offset - possible deviation from the time interval
func (s *Span) Equal(input Span, offset ...time.Duration) bool {
	defaultOffset := time.Second * 0
	if len(offset) > 0 {
		defaultOffset = offset[0]
	}
	startSub := s.start.Sub(input.start)
	endSub := s.end.Sub(input.end)

	return startSub <= defaultOffset && startSub >= -defaultOffset &&
		endSub <= defaultOffset && endSub >= -defaultOffset
}

// IsIntersection  проверка на пересечение временных интервалов (не включаются случаи на стыке)
// offset - possible deviation from the time interval
func (s *Span) IsIntersection(input Span, offset ...time.Duration) bool {
	defaultOffset := time.Second * 0
	if len(offset) > 0 {
		defaultOffset = offset[0]
	}
	return s.start.Add(defaultOffset).Before(input.end) && s.end.After(input.start.Add(defaultOffset))
}
