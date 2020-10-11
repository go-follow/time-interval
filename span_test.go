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
				time.Date(2020, 10, 11, 15, 00, 0, 0, time.UTC),
				time.Date(2020, 10, 11, 16, 00, 0, 0, time.UTC)),
			offset: time.Second * 5,
			excepted: true,
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
			offset: 3 * time.Hour,
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
