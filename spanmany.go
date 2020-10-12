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

// ExceptionIfIntersection excludes periods from the SpanMany if there is an intersection with another SpanMany
func (s *SpanMany) ExceptionIfIntersection(input SpanMany, offset ...time.Duration) SpanMany {
	var listSpans []Span
	for _, s := range s.spans {
		if input.IsIntersection(s, offset...) {
			continue
		}
		listSpans = append(listSpans, s)
	}
	return NewMany(listSpans...)
}

// ExceptionIfNotEqual excludes periods from the SpanMany if it does not meet any equality with another SpanMany
func (s *SpanMany) ExceptionIfNotEqual(input SpanMany, offset ...time.Duration) SpanMany {
	var listSpans []Span
	for _, s := range s.spans {
		if input.Equal(s, offset...) {
			listSpans = append(listSpans, s)
			continue
		}
	}
	return NewMany(listSpans...)
}

// Intersection - нахождение временных интервалов по пересечению с другим временным интервалом
func (s *SpanMany) Intersection(input Span) SpanMany {
	if len(s.spans) == 0 {
		return NewMany()
	}
	var intersectionList []Span
	for _, sp := range s.spans {
		intersectionSpan := sp.Intersection(input)
		if !intersectionSpan.IsEmpty() {
			intersectionList = append(intersectionList, intersectionSpan)
		}
	}
	return NewMany(intersectionList...)
}
