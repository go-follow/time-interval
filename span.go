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
func (s *Span) IsEmpty() bool {
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

// IsIntersection check for intersection of time intervals
// offset - possible deviation from the time interval
func (s *Span) IsIntersection(input Span, offset ...time.Duration) bool {
	defaultOffset := time.Second * 0
	if len(offset) > 0 {
		defaultOffset = offset[0]
	}
	return s.start.Add(defaultOffset).Before(input.end) && s.end.After(input.start.Add(defaultOffset))
}

// Intersection - intersection of two time intervals
func (s *Span) Intersection(input Span) Span {
	if !s.IsIntersection(input) {
		return Span{}
	}
	if afterOrEqual(s.end, input.start) &&
		beforeOrEqual(s.start, input.start) && beforeOrEqual(s.end, input.end) {
		return New(input.start, s.end)
	}
	if beforeOrEqual(s.start, input.end) &&
		afterOrEqual(s.end, input.end) && afterOrEqual(s.start, input.start) {
		return New(s.start, input.end)
	}
	if afterOrEqual(s.start, input.start) && beforeOrEqual(s.end, input.end) {
		return New(s.start, s.end)
	}
	if beforeOrEqual(s.start, input.start) && afterOrEqual(s.end, input.end) {
		return New(input.start, input.end)
	}
	panic("unknown case for Intersection")
}
