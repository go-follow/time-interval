package time_interval

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	testCases := []struct {
		name          string
		newInterval   Span
		inputInterval Span
		offset        time.Duration
		excepted      bool
	}{
		{
			name: "not_equal_slightly",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 30, 17, 12, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 30, 17, 11, time.UTC)),
			excepted: false,
		},
		{
			name: "equal_slightly_with_offset",
			newInterval: New(
				time.Date(2020, 10, 11, 15, 0, 5, 0, time.UTC),
				time.Date(2020, 10, 11, 16, 0, 5, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC)),
			offset:   time.Second * 5,
			excepted: true,
		},
		{
			name: "equal_offset_5_minute",
			newInterval: New(
				time.Date(2020, 10, 11, 15, 5, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 16, 5, 0, 0, time.UTC)),
			offset:   time.Minute * 5,
			excepted: true,
		},
		{
			name: "not_equal_offset_5_minute",
			newInterval: New(
				time.Date(2020, 10, 11, 15, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 16, 20, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 16, 20, 0, 0, time.UTC)),
			offset:   time.Minute * 5,
			excepted: false,
		},
		{
			name: "not_equal_many",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 22, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 23, 30, 17, 11, time.UTC)),
			excepted: false,
		},
		{
			name: "equal_many_with_offset",
			newInterval: New(
				time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 21, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 14, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 0, 0, 0, time.UTC)),
			offset:   3 * time.Hour,
			excepted: true,
		},
		{
			name: "full_equal",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 30, 17, 12, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 30, 17, 12, time.UTC)),
			excepted: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newInterval.Equal(tc.inputInterval, tc.offset)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestIsIntersection(t *testing.T) {
	testCases := []struct {
		name          string
		newInterval   Span
		inputInterval Span
		offset        time.Duration
		excepted      bool
	}{
		{
			name: "invert_case",
			newInterval: New(
				time.Date(2020, 10, 11, 11, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC)),
			excepted: true,
		},
		{
			name: "intersection_slightly",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			),
			inputInterval: New(
				time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 30, 0, 1, time.UTC),
			),
			excepted: true,
		},
		{
			name: "not_intersection_slightly_with_offset",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			),
			inputInterval: New(
				time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 30, 5, 0, time.UTC),
			),
			offset:   time.Second * 5,
			excepted: false,
		},
		{
			name: "intersection_slightly_with_offset",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
			),
			inputInterval: New(
				time.Date(2020, 10, 11, 16, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 30, 6, 0, time.UTC),
			),
			offset:   time.Second * 5,
			excepted: true,
		},
		{
			name: "intersection_many",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 17, 10, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC)),
			excepted: true,
		},
		{
			name: "not_intersection_many",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 17, 10, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC)),
			offset:   time.Hour * 1,
			excepted: false,
		},
		{
			name: "not_intersection",
			newInterval: New(
				time.Date(2020, 10, 11, 17, 30, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 11, 18, 22, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 19, 0, 0, 0, time.UTC)),
			excepted: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newInterval.IsIntersection(tc.inputInterval, tc.offset)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestIntersection(t *testing.T) {
	testCases := []struct {
		name      string
		newSpan   Span
		inputSpan Span

		excepted Span
	}{
		{
			name: "input_contains_new",
			newSpan: New(
				time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 11, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 19, 0, 0, 0, time.UTC),
			),
			excepted: New(
				time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 17, 0, 0, 0, time.UTC),
			),
		},
		{
			name: "new_contains_input",
			newSpan: New(
				time.Date(2020, 10, 11, 11, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 14, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 13, 0, 0, 0, time.UTC),
			),
			excepted: New(
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 13, 0, 0, 0, time.UTC),
			),
		},
		{
			name: "not_intersection",
			newSpan: New(
				time.Date(2020, 10, 11, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
			),
			excepted: Span{},
		},
		{
			name: "not_intersection_many",
			newSpan: New(
				time.Date(2020, 10, 11, 3, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 7, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 11, 22, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 23, 0, 0, 0, time.UTC),
			),
			excepted: Span{},
		},
		{
			name: "intersection_new_left",
			newSpan: New(
				time.Date(2020, 10, 11, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 11, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 15, 0, 0, 0, time.UTC),
			),
			excepted: New(
				time.Date(2020, 10, 11, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			),
		},
		{
			name: "intersection_new_right",
			newSpan: New(
				time.Date(2020, 10, 11, 11, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 14, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 11, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			),
			excepted: New(
				time.Date(2020, 10, 11, 11, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 12, 0, 0, 0, time.UTC),
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newSpan.Intersection(tc.inputSpan)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		name          string
		newInterval   Span
		inputInterval Span

		excepted SpanMany
	}{
		{
			name: "boundaries are equal",
			newInterval: New(
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 9, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 1, 9, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC)),
			excepted: NewMany(New(
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC))),
		},
		{
			name: "new_absorbs_input",
			newInterval: New(
				time.Date(2020, 10, 1, 7, 15, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 17, 23, 11, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 11, 0, 11, 0, time.UTC)),
			excepted: NewMany(New(
				time.Date(2020, 10, 1, 7, 15, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 17, 23, 11, 0, time.UTC))),
		},
		{
			name: "input_absorbs_new",
			newInterval: New(
				time.Date(2020, 10, 1, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 14, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC)),
			excepted: NewMany(New(
				time.Date(2020, 10, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC))),
		},
		{
			name: "new_left_input",
			newInterval: New(
				time.Date(2020, 10, 1, 5, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 12, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 1, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 14, 0, 0, 0, time.UTC)),
			excepted: NewMany(New(
				time.Date(2020, 10, 1, 5, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 14, 0, 0, 0, time.UTC))),
		},
		{
			name: "input_left_new",
			newInterval: New(
				time.Date(2020, 10, 1, 21, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 23, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)),
			excepted: NewMany(New(
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 23, 0, 0, 0, time.UTC))),
		},
		{
			name: "new_next_input",
			newInterval: New(
				time.Date(2020, 10, 1, 4, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)),
			excepted: NewMany(
				New(
					time.Date(2020, 10, 1, 4, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC)),
				New(
					time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC))),
		},
		{
			name: "input_next_new",
			newInterval: New(
				time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)),
			inputInterval: New(
				time.Date(2020, 10, 1, 4, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC)),
			excepted: NewMany(
				New(
					time.Date(2020, 10, 1, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 1, 22, 0, 0, 0, time.UTC)),
				New(
					time.Date(2020, 10, 1, 4, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 1, 7, 0, 0, 0, time.UTC))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newInterval.Union(tc.inputInterval)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestExcept(t *testing.T) {
	testCases := []struct {
		name      string
		newSpan   Span
		inputSpan Span

		excepted SpanMany
	}{
		{
			name: "not_intersection",
			newSpan: New(
				time.Date(2020, 10, 14, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 18, 0, 0, 0, time.UTC),
			),
			excepted: NewMany(
				New(
					time.Date(2020, 10, 14, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				),
			),
		},
		{
			name: "full_intersection",
			newSpan: New(
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 14, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 14, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 15, 0, 0, 0, time.UTC),
			),
			excepted: NewMany(),
		},
		{
			name: "torn_result",
			newSpan: New(
				time.Date(2020, 10, 14, 7, 15, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 15, 22, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 14, 0, 0, 0, time.UTC),
			),
			excepted: NewMany(
				New(
					time.Date(2020, 10, 14, 7, 15, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
				),
				New(
					time.Date(2020, 10, 14, 14, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 15, 22, 0, 0, time.UTC),
				),
			),
		},
		{
			name: "right_takeover",
			newSpan: New(
				time.Date(2020, 10, 14, 9, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 12, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 14, 11, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 15, 0, 0, 0, time.UTC),
			),
			excepted: NewMany(
				New(
					time.Date(2020, 10, 14, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 11, 0, 0, 0, time.UTC),
				),
			),
		},
		{
			name: "left_takeover",
			newSpan: New(
				time.Date(2020, 10, 14, 10, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 15, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 10, 14, 9, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 14, 14, 30, 0, 0, time.UTC),
			),
			excepted: NewMany(
				New(
					time.Date(2020, 10, 14, 14, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 14, 15, 0, 0, 0, time.UTC),
				),
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newSpan.Except(tc.inputSpan)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestIsContains(t *testing.T) {
	testCases := []struct {
		name      string
		newSpan   Span
		inputSpan Span
		offset    time.Duration
		excepted  bool
	}{
		{
			name: "contains",
			newSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 8, 30, 0, 0, time.UTC),
			),
			excepted: true,
		},
		{
			name: "equal",
			newSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			excepted: true,
		},
		{
			name: "reverse_contains",
			newSpan: New(
				time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 8, 30, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			offset: time.Hour,
			excepted: true,
		},
		{
			name: "not_contains_left",
			newSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 11, 20, 6, 59, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			excepted: false,
		},
		{
			name: "contains_left_offset",
			newSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 11, 20, 6, 59, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			offset: time.Minute,
			excepted: true,
		},
		{
			name: "not_contains_right",
			newSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 1, time.UTC),
			),
			excepted: false,
		},
		{
			name: "contains_right_offset",
			newSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 11, 20, 8, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 1, time.UTC),
			),
			offset: time.Second,
			excepted: true,
		},
		{
			name: "many_not_contains",
			newSpan: New(
				time.Date(2020, 11, 20, 7, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 9, 0, 0, 0, time.UTC),
			),
			inputSpan: New(
				time.Date(2020, 11, 20, 1, 0, 0, 0, time.UTC),
				time.Date(2020, 11, 20, 5, 0, 0, 0, time.UTC),
			),
			excepted: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newSpan.IsContains(tc.inputSpan, tc.offset)
			assert.Equal(t, tc.excepted, result)
		})
	}
}
