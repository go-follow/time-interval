package timeinterval

import (
	"fmt"
	"sort"
	"time"
)

// SpanMany  model containing more than one time interval.
type SpanMany struct {
	spans []Span
}

// NewMany initialization for multiple time intervals.
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

// Add adding a time interval to SpanMany.
func (s *SpanMany) Add(start time.Time, end time.Time) error {
	if s == nil || s.spans == nil {
		return nil
	}
	interval, err := New(start, end)
	if err != nil {
		return err
	}
	s.spans = append(s.spans, interval)
	return nil
}

// AddMany adding several time slots at once to the existing one SpanMany.
func (s *SpanMany) AddMany(spans ...Span) {
	if s.spans == nil || len(spans) == 0 {
		return
	}
	s.spans = append(s.spans, spans...)
}

// String implementation interface stringer for SpanMany.
func (s *SpanMany) String() string {
	str := "["
	for _, s := range s.spans {
		str += fmt.Sprintf("\n\t%v - %v", s.Start(), s.End())
	}
	str += "\n]"
	return str
}

// Spans get an array of intervals.
func (s *SpanMany) Spans() []Span {
	if s.spans == nil {
		return []Span{}
	}
	return s.spans
}

// Sort sorting time intervals.

// st - sorting options:
// Ascending sort Ascending (default)
// Descending sort descending
func (s *SpanMany) Sort(st ...SortType) {
	if len(s.spans) == 0 {
		return
	}
	if len(st) > 0 && st[0] == Descending {
		sort.Slice(s.spans, func(i, j int) bool {
			return s.spans[i].start.After(s.spans[j].start)
		})
		return
	}
	sort.Slice(s.spans, func(i, j int) bool {
		return s.spans[i].start.Before(s.spans[j].start)
	})
}

// Equal full comparison of SpanMany of time intervals with one interval.
// If there is at least one match, return true.

// offset - possible deviation from the time interval.
func (s *SpanMany) Equal(input Span, offset ...time.Duration) bool {
	for _, s := range s.spans {
		if s.Equal(input, offset...) {
			return true
		}
	}
	return false
}

// IsIntersection checking for intersection of an interval with one of SpanMany.
// If there is at least one match, return true.

// offset - possible deviation from the time interval.
func (s *SpanMany) IsIntersection(input Span, offset ...time.Duration) bool {
	for _, s := range s.spans {
		if s.IsIntersection(input, offset...) {
			return true
		}
	}
	return false
}

// IsContains checking for contains of SpanMany of time intervals with one interval.
// If there is at least one match, return true.
func (s *SpanMany) IsContains(input Span, offset ...time.Duration) bool {
	for _, s := range s.spans {
		if s.IsContains(input, offset...) {
			return true
		}
	}
	return false
}

// ExceptionIfIntersection excludes periods from the SpanMany if there is an intersection with another SpanMany.

// offset - possible deviation from the time interval
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

// ExceptionIfNotEqual excludes periods from the SpanMany if it does not meet any equality with another SpanMany.

// offset - possible deviation from the time interval
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

// ExceptionIfNotContains excludes periods from SpanMany if it does not contain any interval with another SpanMany.
func (s *SpanMany) ExceptionIfNotContains(input SpanMany, offset ...time.Duration) SpanMany {
	var listSpans []Span
	for _, sp := range s.spans {
		if input.IsContains(sp, offset...) {
			listSpans = append(listSpans, sp)
			continue
		}
	}
	return NewMany(listSpans...)
}

// Intersection intersecting time slots (SpanMany) with one time slot (Span).
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

// Except difference between each array element SpanMany and input Span  (s[i] \ input).
// Returns the elements of SpanMany, where the time interval remains after the Except operation with input.
// Before returning the final result is sorted and merged.
func (s *SpanMany) Except(input Span) SpanMany {
	result := NewMany()
	for _, sp := range s.spans {
		ex := sp.Except(input)
		if len(ex.spans) == 0 {
			continue
		}
		result.AddMany(ex.spans...)
	}
	return result.Union()
}

// Union concatenation SpanMany of array SpanMany.
func (s *SpanMany) Union(input ...SpanMany) SpanMany {
	var result []Span
	for _, inp := range input {
		s.spans = append(s.spans, inp.spans...)
	}
	s.Sort()

	var bufferSpan Span
	for _, sp := range s.spans {
		if bufferSpan.start.IsZero() {
			bufferSpan = sp
			continue
		}
		if sp.isIntersectionEqual(bufferSpan) {
			bufferSpan = Span{
				start: sp.minStart(bufferSpan),
				end:   sp.maxEnd(bufferSpan),
			}
			continue
		}
		result = append(result, bufferSpan)
		bufferSpan = sp
	}
	if !bufferSpan.start.IsZero() {
		result = append(result, bufferSpan)
	}
	return NewMany(result...)
}
