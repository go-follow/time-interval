package time_interval

import "time"

// SpanMany  model containing more than one time interval
type SpanMany struct {
	spans []Span
}

// NewMany  initialization for multiple time intervals
func NewMany(spans ...Span) SpanMany {
	if len(spans) == 0 {
		return SpanMany{
			spans: []Span{},
		}
	}
	return SpanMany{
		spans: spans,
	}
}

// Equal  full comparison of SpanMany of time intervals with one interval
// If there is at least one match, return true
func (s *SpanMany) Equal(input Span, offset ...time.Duration) bool {
	for _, s := range s.spans {
		if s.Equal(input, offset...) {
			return true
		}
	}
	return false
}

// IsIntersection  checking for intersection of an interval with one of SpanMany
// If there is at least one match, return true
func (s *SpanMany) IsIntersection(input Span, offset ...time.Duration) bool {
	if len(s.spans) == 0 {
		return false
	}
	for _, s := range s.spans {
		if s.IsIntersection(input, offset...) {
			return true
		}
	}
	return false
}
