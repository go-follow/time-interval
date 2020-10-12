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

func TestExceptionIfIntersection(t *testing.T) {
	testCases := []struct {
		name              string
		newIntervalMany   SpanMany
		inputIntervalMany SpanMany

		wantResult SpanMany
	}{
		{
			newIntervalMany: NewMany(
				New(time.Date(2020, 10, 12, 7, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 8, 0, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 7, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 8, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC)),
			),
			inputIntervalMany: NewMany(
				New(time.Date(2020, 10, 12, 6, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 7, 35, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 31, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 8, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 8, 0, 0, 1, time.UTC)),
			),

			wantResult: NewMany(
				New(
					time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)),
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			listSpan := tc.newIntervalMany.ExceptionIfIntersection(tc.inputIntervalMany)
			assert.Equal(t, tc.wantResult, listSpan)
		})
	}
}

func TestExceptionIfNotEqual(t *testing.T) {
	testCases := []struct {
		name              string
		newIntervalMany   SpanMany
		inputIntervalMany SpanMany

		wantResult SpanMany
	}{
		{
			newIntervalMany: NewMany(
				New(time.Date(2020, 10, 12, 8, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 9, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 10, 30, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 11, 0, 0, 0, time.UTC)),
			),
			inputIntervalMany: NewMany(
				New(time.Date(2020, 10, 12, 7, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 8, 35, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 29, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)),
				New(time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 11, 0, 0, 1, time.UTC)),
				New(time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 11, 0, 0, 0, time.UTC)),
			),

			wantResult: NewMany(
				New(
					time.Date(2020, 10, 12, 8, 30, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 9, 30, 0, 0, time.UTC)),
				New(
					time.Date(2020, 10, 12, 10, 0, 0, 0, time.UTC),
					time.Date(2020, 10, 12, 11, 0, 0, 0, time.UTC)),
			),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			listSpan := tc.newIntervalMany.ExceptionIfNotEqual(tc.inputIntervalMany)
			assert.Equal(t, tc.wantResult, listSpan)
		})
	}
}
