package timeinterval

import (
	"errors"
	"fmt"
	"time"
)

// Span time interval
type Span struct {
	start time.Time
	end   time.Time
}

// New initialization of a new time interval.

func New(start, end time.Time) (Span, error) {
	if afterOrEqual(start, end) {
		return Span{}, errors.New("time start cannot be more time end")
	}
	return Span{
		start: start,
		end:   end,
	}, nil
}

// Start returning start time interval
func (s *Span) Start() time.Time {
	return s.start
}

// End returning end time interval
func (s *Span) End() time.Time {
	return s.end
}

// String implementation interface stringer for Span
func (s *Span) String() string {
	return fmt.Sprintf("%v - %v", s.Start(), s.End())
}

// IsEmpty  defines empty spacing
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

//IsContains check contains interval
func (s *Span) IsContains(input Span, offset ...time.Duration) bool {
	defaultOffset := time.Second * 0
	if len(offset) > 0 {
		defaultOffset = offset[0]
	}
	return beforeOrEqual(s.start.Add(-defaultOffset), input.start) &&
		afterOrEqual(s.end.Add(defaultOffset), input.end)
}

// Intersection intersection of two time intervals
func (s *Span) Intersection(input Span) Span {
	if !s.IsIntersection(input) {
		return Span{}
	}
	if afterOrEqual(s.end, input.start) &&
		beforeOrEqual(s.start, input.start) && beforeOrEqual(s.end, input.end) {
		return Span{
			start: input.start,
			end:   s.end,
		}
	}
	if beforeOrEqual(s.start, input.end) &&
		afterOrEqual(s.end, input.end) && afterOrEqual(s.start, input.start) {
		return Span{
			start: s.start,
			end:   input.end,
		}
	}
	if afterOrEqual(s.start, input.start) && beforeOrEqual(s.end, input.end) {
		return Span{
			start: s.start,
			end:   s.end,
		}
	}
	if beforeOrEqual(s.start, input.start) && afterOrEqual(s.end, input.end) {
		return Span{
			start: input.start,
			end:   input.end,
		}
	}
	return Span{}
}

// Union union of two time intervals.
func (s *Span) Union(input Span) SpanMany {
	if s.isIntersectionEqual(input) {
		return NewMany(Span{
			start: s.minStart(input),
			end:   s.maxEnd(input),
		})
	}
	result := NewMany(
		Span{
			start: s.start,
			end:   s.end,
		},
		Span{
			start: input.start,
			end:   input.end,
		})
	return result
}

// Except  difference in time intervals - from input (s \ input).
func (s *Span) Except(input Span) SpanMany {
	if !s.IsIntersection(input) {
		return NewMany(Span{
			start: s.start,
			end:   s.end,
		})
	}
	if afterOrEqual(s.start, input.start) && beforeOrEqual(s.end, input.end) {
		return NewMany()
	}
	if s.start.Before(input.start) && s.end.After(input.end) {
		return NewMany(
			Span{
				start: s.start,
				end:   input.start,
			},
			Span{
				start: input.end,
				end:   s.end,
			},
		)
	}
	if s.start.Before(input.start) && afterOrEqual(s.end, input.start) {
		return NewMany(Span{
			start: s.start,
			end:   input.start,
		})
	}
	if beforeOrEqual(s.start, input.end) && s.end.After(input.end) {
		return NewMany(Span{
			start: input.end,
			end:   s.end,
		})
	}
	return SpanMany{}
}

func (s *Span) minStart(input Span) time.Time {
	if s.start.Before(input.start) {
		return s.start
	}
	return input.start
}

func (s *Span) maxEnd(input Span) time.Time {
	if s.end.After(input.end) {
		return s.end
	}
	return input.end
}

// isIntersectionEqual checking for the intersection of time intervals.
// The difference from the public method is that it includes cases at the junction.
func (s *Span) isIntersectionEqual(input Span) bool {
	return beforeOrEqual(s.start, input.end) && afterOrEqual(s.end, input.start)
}
