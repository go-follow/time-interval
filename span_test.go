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
