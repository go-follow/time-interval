package time_interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEqualMany(t *testing.T) {
	testCases := []struct {
		name            string
		newIntervalMany SpanMany
		inputInterval   Span
		excepted        bool
	}{
		{
			name: "not_equal_slightly",
			newIntervalMany: NewMany(
				New(
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 12, time.UTC)),
				New(
					time.Date(2020, 10, 12, 17, 0, 0, 1, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
				New(
					time.Date(2020, 10, 12, 17, 0, 0, 1, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 12, time.UTC)),
			),
			inputInterval: New(
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
			excepted: false,
		},
		{
			name: "not_equal_many",
			newIntervalMany: NewMany(
				New(
					time.Date(2020, 10, 12, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 14, 0, 0, 0, time.UTC)),
				New(
					time.Date(2020, 10, 12, 15, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 16, 0, 0, 0, time.UTC)),
				New(
					time.Date(2020, 10, 12, 21, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 22, 0, 0, 0, time.UTC)),
			),
			inputInterval: New(
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
			excepted: false,
		},
		{
			name: "equal",
			newIntervalMany: NewMany(
				New(
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 12, time.UTC)),
				New(
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
				New(
					time.Date(2020, 10, 12, 19, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC)),
				New(
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
			),
			inputInterval: New(
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
			excepted: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newIntervalMany.Equal(tc.inputInterval)
			assert.Equal(t, tc.excepted, result)
		})
	}
}

func TestIsIntersectionMany(t *testing.T) {
	testCases := []struct {
		name            string
		newIntervalMany SpanMany
		inputInterval   Span
		excepted        bool
	}{
		{
			name: "not_intersection_slightly",
			newIntervalMany: NewMany(
				New(
					time.Date(2020, 10, 12, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 13, 0, 0, 12, time.UTC)),
				New(
					time.Date(2020, 10, 12, 17, 0, 0, 1, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
				New(
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC),
					time.Date(2020, 10, 12, 110, 0, 0, 0, time.UTC)),
			),
			inputInterval: New(
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
			excepted: false,
		},
		{
			name: "not_intersection_many",
			newIntervalMany: NewMany(
				New(
					time.Date(2020, 10, 12, 12, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 13, 0, 0, 12, time.UTC)),
				New(
					time.Date(2020, 10, 12, 22, 0, 0, 1, time.UTC),
					time.Date(2020, 10, 12, 23, 0, 0, 11, time.UTC)),
				New(
					time.Date(2020, 10, 12, 19, 0, 0, 11, time.UTC),
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC)),
			),
			inputInterval: New(
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
			excepted: false,
		},
		{
			name: "equal",
			newIntervalMany: NewMany(
				New(
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 12, time.UTC)),
				New(
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 10, time.UTC)),
				New(
					time.Date(2020, 10, 12, 19, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 20, 0, 0, 0, time.UTC)),
				New(
					time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
			),
			inputInterval: New(
				time.Date(2020, 10, 12, 17, 0, 0, 0, time.UTC),
				time.Date(2020, 10, 12, 18, 0, 0, 11, time.UTC)),
			excepted: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.newIntervalMany.Equal(tc.inputInterval)
			assert.Equal(t, tc.excepted, result)
		})
	}
}
